package iso8583

import (
	lib8583 "github.com/mofax/iso8583"
)

// MessageFactory to create instances of the ISO8583 messages
type MessageFactory struct {
	Specification lib8583.IsoStruct
}

// DefaultMessageFactory - Creates a default instance of the MessageFactory
func DefaultMessageFactory(specFile string) *MessageFactory {
	return &MessageFactory{lib8583.NewISOStruct(specFile, true)}
}
