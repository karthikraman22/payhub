package client

import (
	"encoding/binary"
	"io"
	"net"

	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
	log "github.com/sirupsen/logrus"
)

// ClientHandler to handle the ISO8583 messages
type ClientHandler struct {
	Encoder *iso8583.Encoder
	Decoder *iso8583.Decoder
}

func recovery1() {
	if r := recover(); r != nil {
		log.Errorf("recovered from error %v, closing server connection", r)
	}
}

// Handle - Implements the handling of a new iso8583 client
func (ch *ClientHandler) Handle(connection *net.TCPConn) {
	defer recovery1()
	defer connection.Close()
	mli := make([]byte, ch.Decoder.HeaderLength)
	for {
		n, err := connection.Read(mli)
		if err == io.EOF {
			log.Print("eof received")
			// Client connection closed
			break
		}
		reqLen := binary.BigEndian.Uint16(mli)
		log.Printf("reading incoming message with %d bytes", reqLen)
		//all good, read rest of the message data
		msgData := make([]byte, reqLen)
		n, err = connection.Read(msgData)
		if err != nil {
			log.Errorf("error while reading data from client %s,  error code %v", connection.RemoteAddr().String(), err)
			break
		}
		if uint16(n) != reqLen {
			log.Printf("not enough data - required: %d != actual %d", reqLen, n)
			continue
		}
		go ch.handleIncomingMsg(msgData, connection)
	}
	log.Printf("closing connection @ %s", connection.RemoteAddr().String())
}

func (ch *ClientHandler) handleIncomingMsg(msgData []byte, connection *net.TCPConn) {
	isoMsg, err := ch.Decoder.Decode(msgData)
	if err != nil {
		log.Errorf("decoding failed : %v", err)
		return
	}
	log.Printf("incoming message from @ %s :  %v", connection.RemoteAddr().String(), isoMsg.Elements)

}

func (ch *ClientHandler) handleOutgoingMsg(isoMsg *lib8583.IsoStruct, connection *net.TCPConn) error {
	packedMsg, err := ch.Encoder.Encode(isoMsg)
	if err != nil {
		log.Errorf("unable to pack message %v", err)
		return err
	}
	log.Printf("outgoing message to @ %s :  %v", connection.RemoteAddr().String(), isoMsg.Elements)
	_, err = connection.Write(packedMsg)
	return err
}
