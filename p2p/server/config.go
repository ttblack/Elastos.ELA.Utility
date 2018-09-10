package server

import (
	"net"
	"time"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
	"github.com/elastos/Elastos.ELA.Utility/p2p/peer"
)

const (
	defaultMaxPeers              = 125
	defaultBanThreshold   uint32 = 100
	defaultBanDuration           = time.Hour * 24
	defaultConnectTimeout        = time.Second * 30
)

// Config is a descriptor which specifies the server instance configuration.
type Config struct {
	MagicNumber      uint32
	SeedPeers        []string
	ListenAddrs      []string
	ExternalIPs      []string
	Upnp             bool
	DefaultPort      uint16
	DisableListen    bool
	DisableRelayTx   bool
	MaxPeers         int
	DisableBanning   bool
	BanThreshold     uint32
	BanDuration      time.Duration
	Whitelists       []*net.IPNet
	TargetOutbound   int
	OnNewPeer        func(*peer.Peer)
	OnDonePeer       func(*peer.Peer)
	MakeEmptyMessage func(string) (p2p.Message, error)
	BestHeight       func() uint64
}

// inWhitelist returns whether the IP address is included in the whitelisted
// networks and IPs.
func (cfg *Config) inWhitelist(addr net.Addr) bool {
	if len(cfg.Whitelists) == 0 {
		return false
	}

	host, _, err := net.SplitHostPort(addr.String())
	if err != nil {
		log.Warnf("Unable to SplitHostPort on '%s': %v", addr, err)
		return false
	}
	ip := net.ParseIP(host)
	if ip == nil {
		log.Warnf("Unable to parse IP '%s'", addr)
		return false
	}

	for _, ipnet := range cfg.Whitelists {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

func dialTimeout(addr net.Addr) (net.Conn, error) {
	return net.DialTimeout(addr.Network(), addr.String(), defaultConnectTimeout)
}

// NewDefaultConfig returns a new config instance filled by default settings
// for the server.
func NewDefaultConfig(
	magic uint32,
	defaultPort uint16,
	seeds, listenAddrs []string,
	onNewPeer func(*peer.Peer),
	onDonePeer func(*peer.Peer),
	makeEmptyMessage func(string) (p2p.Message, error),
	bestHeight func() uint64) *Config {
	return &Config{
		MagicNumber:      magic,
		SeedPeers:        seeds,
		ListenAddrs:      listenAddrs,
		ExternalIPs:      nil,
		Upnp:             false,
		DefaultPort:      defaultPort,
		DisableListen:    false,
		DisableRelayTx:   false,
		MaxPeers:         defaultMaxPeers,
		DisableBanning:   false,
		BanThreshold:     defaultBanThreshold,
		BanDuration:      defaultBanDuration,
		Whitelists:       nil,
		TargetOutbound:   defaultTargetOutbound,
		OnNewPeer:        onNewPeer,
		OnDonePeer:       onDonePeer,
		MakeEmptyMessage: makeEmptyMessage,
		BestHeight:       bestHeight,
	}
}
