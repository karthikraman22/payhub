package handlers

import (
	"errors"

	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
)

// IsoMsgHandler -
type IsoMsgHandler interface {
	Handle(rq *lib8583.IsoStruct) (*lib8583.IsoStruct, error)
}

// GetHandler - creates and returns a new message handler for the given MTI
func GetHandler(mti string, mf *iso8583.MessageFactory) (IsoMsgHandler, error) {
	switch mti {
	case "0200":
		return AuthRqHandler{MessageFactory: mf}, nil
	case "0400":
		return ReversalRqHandler{MessageFactory: mf}, nil
	default:
		return nil, errors.New("Invalid mti - " + mti)
	}
}
