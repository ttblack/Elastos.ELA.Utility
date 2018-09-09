package msg

import (
	"github.com/elastos/Elastos.ELA.Utility/common"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

const MaxBlockSize = 8000000

type Block struct {
	common.Serializable
}

func NewBlock(block common.Serializable) *Block {
	return &Block{Serializable: block}
}

func (msg *Block) CMD() string {
	return p2p.CmdBlock
}

func (msg *Block) MaxLength() uint32 {
	return MaxBlockSize
}

func (msg *Block) Serialize(writer io.Writer) error {
	return msg.Serializable.Serialize(writer)
}

func (msg *Block) Deserialize(reader io.Reader) error {
	return msg.Serializable.Deserialize(reader)
}
