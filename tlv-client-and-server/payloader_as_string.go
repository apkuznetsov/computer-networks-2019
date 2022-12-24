package tlv

import (
	"encoding/binary"
	"errors"
	"io"
)

type String string

func (m String) Bytes() []byte { return []byte(m) }

func (m String) String() string { return string(m) }

func (m String) WriteTo(w io.Writer) (int64, error) {
	// write type
	err := binary.Write(w, binary.BigEndian, StringType) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	// write payload size
	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4

	// write payload
	o, err := w.Write([]byte(m)) // payload

	return n + int64(o), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	// read type
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	if typ != StringType {
		return n, errors.New("invalid String")
	}

	// read size
	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4

	// read payload
	buf := make([]byte, size)
	o, err := r.Read(buf) // payload
	if err != nil {
		return n, err
	}
	*m = String(buf)

	return n + int64(o), nil
}
