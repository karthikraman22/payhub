package handlers

import (
	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
)

func handleAuthRq(rq *lib8583.IsoStruct, mf *iso8583.MessageFactory) (*lib8583.IsoStruct, error) {
	rs := mf.NewInstance("1814", true)
	return rs, nil
}
