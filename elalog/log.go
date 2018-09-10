package elalog

import (
	"log"
	"os"
)

func init() {
	Disabled = &disableLog{}
	Stdout = &stdout{
		logger: log.New(os.Stdout, "S ", log.LstdFlags|log.Lmicroseconds),
	}
}