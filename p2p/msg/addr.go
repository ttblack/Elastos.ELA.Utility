package msg

import (
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

const MaxAddrPerMsg = 1000

// Ensure Addr implement p2p.Message interface.
var _ p2p.Message = (*Addr)(nil)

type Addr struct {
	AddrList []*p2p.NetAddress
}

func NewAddr(addresses []*p2p.NetAddress) *Addr {
	msg := new(Addr)
	msg.AddrList = addresses
	return msg
}

func (msg *Addr) CMD() string {
	return p2p.CmdAddr
}

func (msg *Addr) MaxLength() uint32 {
	return 8 + (MaxAddrPerMsg * 42)
}

func (msg *Addr) Serialize(writer io.Writer) error {
	if err := common.WriteUint64(writer, uint64(len(msg.AddrList))); err != nil {
		return err
	}

	for i := range msg.AddrList {
		if err := msg.AddrList[i].Serialize(writer); err != nil {
			return err
		}
	}
	return nil
}

func (msg *Addr) Deserialize(reader io.Reader) error {
	count, err := common.ReadUint64(reader)
	if err != nil {
		return err
	}

	msg.AddrList = make([]*p2p.NetAddress, count)
	for i := range msg.AddrList {
		msg.AddrList[i] = new(p2p.NetAddress)
		if err := msg.AddrList[i].Deserialize(reader); err != nil {
			return err
		}
	}
	return nil
}
