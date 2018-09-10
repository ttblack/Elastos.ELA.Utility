package msg

import (
	"encoding/binary"
	"io"
	"time"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

// Ensure Version implement p2p.Message interface.
var _ p2p.Message = (*Version)(nil)

type Version struct {
	Version   uint32
	Services  uint64
	TimeStamp uint32
	Port      uint16
	Nonce     uint64
	Height    uint64
	Relay     uint8
}

func (msg *Version) CMD() string {
	return p2p.CmdVersion
}

func (msg *Version) MaxLength() uint32 {
	return 35
}

func (msg *Version) Serialize(writer io.Writer) error {
	return binary.Write(writer, binary.LittleEndian, msg)
}

func (msg *Version) Deserialize(reader io.Reader) error {
	return binary.Read(reader, binary.LittleEndian, msg)
}

func NewVersion(pver uint32, services p2p.ServiceFlag, nonce, height uint64, disableRelayTx bool) *Version {
	var relay uint8
	if !disableRelayTx {
		relay = 1
	}
	return &Version{
		Version:   pver,
		Services:  uint64(services),
		TimeStamp: uint32(time.Now().Unix()),
		Nonce:     nonce,
		Height:    height,
		Relay:     relay,
	}
}
