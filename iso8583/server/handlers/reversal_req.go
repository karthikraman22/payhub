package handlers

import (
	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
)

// ReversalRqHandler - 04xx handler
type ReversalRqHandler struct {
	MessageFactory *iso8583.MessageFactory
}

// Handle -
func (mh ReversalRqHandler) Handle(rq *lib8583.IsoStruct) (*lib8583.IsoStruct, error) {
	rs := mh.MessageFactory.NewInstance("0410", true)
	return rs, nil
}
