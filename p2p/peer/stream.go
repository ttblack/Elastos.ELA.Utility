package peer

import (
	"bytes"
	"fmt"
	"io"

	"github.com/elastos/Elastos.ELA.Utility/p2p"
)

type rw struct {
	// The P2P network ID
	magic uint32

	// Make an empty message instance
	makeEmptyMessage func(command string) (p2p.Message, error)
}

func (rw *rw) ReadMessage(r io.Reader) (p2p.Message, error) {
	// Read message header
	var headerBytes [p2p.HeaderSize]byte
	if _, err := io.ReadFull(r, headerBytes[:]); err != nil {
		return nil, err
	}

	// Deserialize message header
	var hdr p2p.Header
	if err := hdr.Deserialize(headerBytes[:]); err != nil {
		return nil, p2p.ErrInvalidHeader
	}

	// Check for messages from wrong network
	if hdr.Magic != rw.magic {
		return nil, p2p.ErrUnmatchedMagic
	}

	// Create struct of appropriate message type based on the command.
	msg, err := rw.makeEmptyMessage(hdr.GetCMD())
	if err != nil {
		return nil, err
	}

	// Check for message length
	if hdr.Length > msg.MaxLength() {
		return nil, p2p.ErrMsgSizeExceeded
	}

	// Read payload
	payload := make([]byte, hdr.Length)
	_, err = io.ReadFull(r, payload[:])
	if err != nil {
		return nil, err
	}

	// Verify checksum
	if err := hdr.Verify(payload); err != nil {
		return nil, p2p.ErrInvalidPayload
	}

	// Deserialize message
	if err := msg.Deserialize(bytes.NewBuffer(payload)); err != nil {
		return nil, fmt.Errorf("deserialize message %s failed %s", msg.CMD(), err.Error())
	}

	return msg, nil
}

func (rw *rw) WriteMessage(w io.Writer, msg p2p.Message) error {
	// Serialize message
	buf := new(bytes.Buffer)
	if err := msg.Serialize(buf); err != nil {
		return fmt.Errorf("serialize message failed %s", err.Error())
	}
	payload := buf.Bytes()

	// Enforce maximum overall message payload.
	if len(payload) > p2p.MaxMessagePayload {
		return p2p.ErrMsgSizeExceeded
	}

	// Create message header
	hdr, err := p2p.BuildHeader(rw.magic, msg.CMD(), payload).Serialize()
	if err != nil {
		return fmt.Errorf("serialize message header failed %s", err.Error())
	}

	// Write header
	if _, err = w.Write(hdr); err != nil {
		return err
	}

	// Write payload
	_, err = w.Write(payload)
	return err
}
