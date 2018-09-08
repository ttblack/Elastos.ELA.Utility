package msg

import "github.com/elastos/Elastos.ELA.Utility/p2p"

type Pong struct {
	Ping
}

func NewPong(nonce uint64) *Pong {
	pong := new(Pong)
	pong.Nonce = nonce
	return pong
}

func (msg *Pong) CMD() string {
	return p2p.CmdPong
}
