package handlers

import (
	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
)

// NetworkRqHandler - 0200 handler
type NetworkRqHandler struct {
	MessageFactory *iso8583.MessageFactory
}

// Handle -
func (mh *NetworkRqHandler) Handle(rq *lib8583.IsoStruct) (*lib8583.IsoStruct, error) {
	rs := mh.MessageFactory.NewInstance("1814", true)
	return rs, nil
}
