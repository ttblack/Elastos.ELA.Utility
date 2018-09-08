package p2p

import (
	"sync"
	"math/rand"
	"time"

	"github.com/elastos/Elastos.ELA.Utility/p2p/addrmgr"
	"github.com/elastos/Elastos.ELA.Utility/p2p/connmgr"
	"github.com/elastos/Elastos.ELA.Utility/p2p/peer"
	"github.com/btcsuite/btcd/wire"
	"github.com/elastos/daca_course/sourcecode/node/p2p"
)

// broadcastMsg provides the ability to house a bitcoin message to be broadcast
// to all connected peers except specified excluded peers.
type broadcastMsg struct {
	message      Message
	excludePeers []*serverPeer
}

// server provides a bitcoin server for handling communications to and from
// bitcoin peers.
type server struct {
	// The following variables must only be used atomically.
	// Putting the uint64s first makes them 64-bit aligned for 32-bit systems.
	bytesReceived uint64 // Total bytes received from all peers since start.
	bytesSent     uint64 // Total bytes sent by all peers since start.
	started       int32
	shutdown      int32
	shutdownSched int32
	startupTime   int64

	addrManager *addrmgr.AddrManager
	connManager *connmgr.ConnManager
	newPeers    chan *serverPeer
	donePeers   chan *serverPeer
	banPeers    chan *serverPeer
	query       chan interface{}
	broadcast   chan broadcastMsg
	wg          sync.WaitGroup
	quit        chan struct{}
	nat         NAT
	services    ServiceFlag
}

// serverPeer extends the peer to maintain state shared by the server.
type serverPeer struct {
	*peer.Peer

	connReq        *connmgr.ConnReq
	server         *server
	persistent     bool
	relayMtx       sync.Mutex
	disableRelayTx bool
	sentAddrs      bool
	knownAddresses map[string]struct{}
	banScore       connmgr.DynamicBanScore
	quit           chan struct{}
}

// newServerPeer returns a new serverPeer instance. The peer needs to be set by
// the caller.
func newServerPeer(s *server, isPersistent bool) *serverPeer {
	return &serverPeer{
		server:         s,
		persistent:     isPersistent,
		knownAddresses: make(map[string]struct{}),
		quit:           make(chan struct{}),
	}
}


// addKnownAddresses adds the given addresses to the set of known addresses to
// the peer to prevent sending duplicate addresses.
func (sp *serverPeer) addKnownAddresses(addresses []*NetAddress) {
	for _, na := range addresses {
		sp.knownAddresses[addrmgr.NetAddressKey(na)] = struct{}{}
	}
}

// addressKnown true if the given address is already known to the peer.
func (sp *serverPeer) addressKnown(na *NetAddress) bool {
	_, exists := sp.knownAddresses[addrmgr.NetAddressKey(na)]
	return exists
}

// setDisableRelayTx toggles relaying of transactions for the given peer.
// It is safe for concurrent access.
func (sp *serverPeer) setDisableRelayTx(disable bool) {
	sp.relayMtx.Lock()
	sp.disableRelayTx = disable
	sp.relayMtx.Unlock()
}

// relayTxDisabled returns whether or not relaying of transactions for the given
// peer is disabled.
// It is safe for concurrent access.
func (sp *serverPeer) relayTxDisabled() bool {
	sp.relayMtx.Lock()
	isDisabled := sp.disableRelayTx
	sp.relayMtx.Unlock()

	return isDisabled
}

// pushAddrMsg sends an addr message to the connected peer using the provided
// addresses.
func (sp *serverPeer) pushAddrMsg(addresses []*NetAddress) {
	// Filter addresses already known to the peer.
	addrs := make([]*NetAddress, 0, len(addresses))
	for _, addr := range addresses {
		if !sp.addressKnown(addr) {
			addrs = append(addrs, addr)
		}
	}
	known, err := sp.PushAddrMsg(addrs)
	if err != nil {
		log.Errorf("Can't push address message to %s: %v", sp.Peer, err)
		sp.Disconnect()
		return
	}
	sp.addKnownAddresses(known)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}