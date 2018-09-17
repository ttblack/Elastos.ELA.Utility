package msg

import (
	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

// Ensure GetAddr implement p2p.Message interface.
var _ p2p.Message = (*GetAddr)(nil)

type GetAddr struct{ empty }

func (msg *GetAddr) CMD() string {
	return p2p.CmdGetAddr
}

func (msg *GetAddr) MaxLength() uint32 {
	return 0
}

func NewGetAddr() *GetAddr {return &GetAddr{}}