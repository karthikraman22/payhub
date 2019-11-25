package handlers

import (
	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
)

// AuthRqHandler - 0200 handler
type AuthRqHandler struct {
	MessageFactory *iso8583.MessageFactory
}

// Handle -
func (mh AuthRqHandler) Handle(rq *lib8583.IsoStruct) (*lib8583.IsoStruct, error) {
	rs := mh.MessageFactory.NewInstance("0210", true)
	rs.AddField(2, rq.Elements.GetElements()[2])
	return rs, nil
}
