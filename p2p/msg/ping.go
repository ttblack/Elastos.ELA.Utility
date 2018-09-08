package msg

import (
	"encoding/binary"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

type Ping struct {
	Nonce uint64
}

func NewPing(nonce uint64) *Ping {
	ping := new(Ping)
	ping.Nonce = nonce
	return ping
}

func (msg *Ping) CMD() string {
	return p2p.CmdPing
}

func (msg *Ping) MaxLength() uint32 {
	return 8
}

func (msg *Ping) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, msg.Nonce)
}

func (msg *Ping) Deserialize(reader io.Reader) error {
	return binary.Read(reader, binary.LittleEndian, &msg.Nonce)
}
