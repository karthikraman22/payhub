package client

import (
	"net"

	"achuala.in/payhub/iso8583"
	lib8583 "github.com/mofax/iso8583"
	log "github.com/sirupsen/logrus"
)

// Client - Iso client
type IsoClient struct {
	connection *net.TCPConn
	handler    *ClientHandler
	specFile   string
}

// Connect - Creates a new connection to the iso server
func Connect(host, port, specFile string) (*IsoClient, error) {
	serverAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return nil, err
	}
	connection, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		return nil, err
	}
	encoder := &iso8583.Encoder{HeaderLength: 2, ExlcudeHeaderLength: true}
	decoder := &iso8583.Decoder{HeaderLength: 2, MessageFactory: iso8583.DefaultMessageFactory(specFile), ExlcudeHeaderLength: true}
	ch := &ClientHandler{Encoder: encoder, Decoder: decoder}
	go ch.Handle(connection)
	return &IsoClient{connection: connection, handler: ch, specFile: specFile}, nil
}

// Close - closes the underlying connection
func (c *IsoClient) Close() {
	log.Print("closing client connection")
	c.connection.Close()
}

// Send - Sends a ISO message to the server.
func (c *IsoClient) Send(isoMsg *lib8583.IsoStruct) error {
	return c.handler.handleOutgoingMsg(isoMsg, c.connection)
}
