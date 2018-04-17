package transaction

import (
	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA.Utility/common/serialize"
	"io"
)

type UTXOUnspent struct {
	Txid  common.Uint256
	Index uint32
	Value common.Fixed64
}

func (uu *UTXOUnspent) Serialize(w io.Writer) {
	uu.Txid.Serialize(w)
	serialize.WriteUint32(w, uu.Index)
	uu.Value.Serialize(w)
}

func (uu *UTXOUnspent) Deserialize(r io.Reader) error {
	uu.Txid.Deserialize(r)

	index, err := serialize.ReadUint32(r)
	uu.Index = uint32(index)
	if err != nil {
		return err
	}

	uu.Value.Deserialize(r)

	return nil
}