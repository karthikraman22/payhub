package server

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"achuala.in/payhub/iso8583"
	"achuala.in/payhub/iso8583/server/handlers"
	log "github.com/sirupsen/logrus"
)

// Handler - Server handler to handle the ISO8583 messages
type Handler struct {
	Encoder *iso8583.Encoder
	Decoder *iso8583.Decoder
}

func recovery() {
	if r := recover(); r != nil {
		log.Errorf("recovered from error %v, closing client connection", r)
	}
}

// Handle - Implements the handling of a new iso8583 client
func (h *Handler) Handle(connection *net.TCPConn) {
	defer recovery()
	defer connection.Close()
	echoTimer := time.Tick(time.Second * 10)
	stopChannel := make(chan int, 1)
	go h.sendEcho(echoTimer, stopChannel, connection)
	mli := make([]byte, h.Decoder.HeaderLength)
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
		go h.handleIncomingMsg(msgData, connection)
	}
	stopChannel <- 0

	log.Printf("closing connection @ %s", connection.RemoteAddr().String())
}

func (h *Handler) handleIncomingMsg(msgData []byte, connection *net.TCPConn) {
	rqMsg, err := h.Decoder.Decode(msgData)
	if err != nil {
		log.Errorf("decoding failed : %v", err)
		return
	}
	log.Printf("incoming message from @ %s :  %v", connection.RemoteAddr().String(), rqMsg.Elements)
	// Use the message header to get the handler
	rqHandler, err := handlers.GetHandler(rqMsg.Mti.String(), h.Decoder.MessageFactory)
	if err != nil {
		log.Errorf("unable to create a handler : %v", err)
		return
	}
	rsMsg, err := rqHandler.Handle(rqMsg)
	if err != nil {
		log.Errorf("processing request failed : %v", err)
		return
	}
	packedMsg, err := h.Encoder.Encode(rsMsg)
	if err != nil {
		log.Errorf("unable to pack message %v", err)
		return
	}
	_, err = connection.Write(packedMsg)
	if err != nil {
		log.Errorf("unable to send message %v", err)
	}
}

func (h *Handler) sendEcho(tick <-chan time.Time, quit chan int, connection *net.TCPConn) {
	for {
		select {
		case t := <-tick:
			{
				echoMsg := h.Decoder.MessageFactory.NewInstance("1804", true)
				echoMsg.AddField(7, t.Format("MMDDhhmmss"))
				echoMsg.AddField(7, fmt.Sprintf("%02d%02d%02d%02d%02d",
					t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second()))
				echoMsg.AddField(11, fmt.Sprintf("%06d", rand.Int63n(1e16)))
				echoMsg.AddField(12, fmt.Sprintf("%02d%02d%02d",
					t.Hour(), t.Minute(), t.Second()))
				echoMsg.AddField(24, "007")
				echoMsg.AddField(37, fmt.Sprintf("%012d", rand.Int63n(1e16)))
				echoMsg.AddField(53, fmt.Sprintf("%016d", rand.Int63n(1e16)))
				echoMsg.AddField(93, fmt.Sprintf("%05d", rand.Int63n(1e16)))
				echoMsg.AddField(94, fmt.Sprintf("%07d", rand.Int63n(1e16)))
				packedMsg, err := h.Encoder.Encode(echoMsg)
				if err != nil {
					log.Errorf("unable to pack message %v", err)
					return
				}
				_, err = connection.Write(packedMsg)
				if err != nil {
					log.Errorf("unable to send message %v", err)
				}
			}
		case <-quit:
			log.Info("quit received, exiting echo loop")
			return
		}
	}
}
