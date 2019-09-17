package iso8583

import (
	"net"

	"achuala.in/payhub/server"
)

// IsoServer - ISO 8583 host server
type IsoServer struct {
	Host     string
	Port     string
	SpecFile string
}

// Start - Start and listen for incoming connections
func (s *IsoServer) Start() error {
	listenAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(s.Host, s.Port))
	if err != nil {
		return err
	}
	encoder := &Encoder{2}
	decoder := &Decoder{HeaderLength: 2, MessageFactory: DefaultMessageFactory(s.SpecFile)}
	ch := &ClientHandler{Encoder: encoder, Decoder: decoder}
	tcpServer := &server.TCPServer{ListenAddress: listenAddr, ConnHandler: ch}
	err = tcpServer.Start()
	return err
}
