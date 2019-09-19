package iso8583

import (
	"bytes"
	"encoding/binary"

	lib8583 "github.com/mofax/iso8583"
)

// Encoder  - Message encoder model
type Encoder struct {
	HeaderLength       uint32
	ExlcudeHeaderLenth bool
}

// Encode - Encodes the ISO message into ahex bytes
func (enc *Encoder) Encode(msg *lib8583.IsoStruct) ([]byte, error) {
	packedData, err := msg.ToString()
	if err != nil {
		return nil, err
	}
	var mli = make([]byte, 2)
	if enc.ExlcudeHeaderLenth {
		binary.BigEndian.PutUint16(mli, uint16(len(packedData)))
	} else {
		binary.BigEndian.PutUint16(mli, uint16(len(packedData)+2))
	}
	buf := bytes.NewBuffer(mli)
	buf.Write([]byte(packedData))
	return (buf.Bytes()), nil
}
