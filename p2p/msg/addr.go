package msg

import (
	"encoding/binary"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

const MaxAddrPerMsg = 1000

type Addr struct {
	AddrList []p2p.NetAddress
}

func NewAddr(addresses []p2p.NetAddress) *Addr {
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
	return common.WriteElements(writer, uint64(len(msg.AddrList)), msg.AddrList)
}

func (msg *Addr) Deserialize(reader io.Reader) error {
	count, err := common.ReadUint64(reader)
	if err != nil {
		return err
	}

	msg.AddrList = make([]p2p.NetAddress, count)
	return binary.Read(reader, binary.LittleEndian, &msg.AddrList)
}
