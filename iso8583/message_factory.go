package iso8583

import (
	lib8583 "github.com/mofax/iso8583"
)

// MessageFactory to create instances of the ISO8583 messages
type MessageFactory struct {
	Specification string
}

// DefaultMessageFactory - Creates a default instance of the MessageFactory
func DefaultMessageFactory(specFile string) *MessageFactory {
	return &MessageFactory{Specification: specFile}
}

// NewInstance - Creates a new instance
func (mf *MessageFactory) NewInstance(mti string, secondaryBitmap bool) *lib8583.IsoStruct {
	isoMsg := lib8583.NewISOStruct(mf.Specification, secondaryBitmap)
	isoMsg.AddMTI(mti)
	return &isoMsg
}

// DefaultInstance - Creates a default instance
func (mf *MessageFactory) DefaultInstance() *lib8583.IsoStruct {
	isoMsg := lib8583.NewISOStruct(mf.Specification, true)
	return &isoMsg
}
