package msg

// Inventory is the same to Inv message.
type Inventory struct {
	Inv
}

func NewInventory() *Inventory {
	msg := &Inventory{
		Inv: *NewInv(),
	}
	return msg
}
