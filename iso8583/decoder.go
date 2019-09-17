package iso8583

import (
	lib8583 "github.com/mofax/iso8583"
)

// Decoder - ISO 8583 message decoder
type Decoder struct {
	HeaderLength   uint32
	MessageFactory *MessageFactory
}

// Decode - Decodes the incoming bytes into an ISO 8583 message
func (dec *Decoder) Decode(data []byte) (*lib8583.IsoStruct, error) {

	isoMsg, err := dec.MessageFactory.Specification.Parse(string(data))
	if err != nil {
		return nil, err
	}
	return &isoMsg, nil
}
