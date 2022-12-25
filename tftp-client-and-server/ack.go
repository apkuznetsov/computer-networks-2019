package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Ack uint16

func (a Ack) MarshalBinary() ([]byte, error) {
	capacity := 2 + 2 // operation code + block number
	buf := new(bytes.Buffer)
	buf.Grow(capacity)

	err := binary.Write(buf, binary.BigEndian, OpAck) // write operation code
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, a) // write block number
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (a *Ack) UnmarshalBinary(p []byte) error {
	var code OpCode
	r := bytes.NewReader(p)

	err := binary.Read(r, binary.BigEndian, &code) // read operation code
	if err != nil {
		return err
	}

	if code != OpAck {
		return errors.New("invalid ACK")
	}

	return binary.Read(r, binary.BigEndian, a) // read block number
}
