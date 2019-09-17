package iso8583

import (
	"encoding/binary"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

// ClientHandler to handle the ISO8583 messages
type ClientHandler struct {
	Encoder *Encoder
	Decoder *Decoder
}

func recovery() {
	if r := recover(); r != nil {
		log.Errorf("recovered from error %v, closing client connection", r)
	}
}

// Handle - Implements the handling of a new iso8583 client
func (ch *ClientHandler) Handle(connection *net.TCPConn) {
	defer recovery()
	defer connection.Close()
	mli := make([]byte, 2)
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
	log.Printf("incoming message from @ %s :  %v", connection.RemoteAddr().String(), isoMsg)

}
