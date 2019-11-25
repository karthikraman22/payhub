package server

import (
	"net"

	"achuala.in/payhub/iso8583"
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
	encoder := &iso8583.Encoder{HeaderLength: 2, ExlcudeHeaderLength: true}
	decoder := &iso8583.Decoder{HeaderLength: 2, MessageFactory: iso8583.DefaultMessageFactory(s.SpecFile), ExlcudeHeaderLength: true}
	handler := &Handler{Encoder: encoder, Decoder: decoder}
	tcpServer := &server.TCPServer{ListenAddress: listenAddr, ConnHandler: handler}
	err = tcpServer.Start()
	return err
}
