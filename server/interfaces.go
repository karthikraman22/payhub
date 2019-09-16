package server

import "net"

// ConnectionHandler interface to handle the client messages
type ConnectionHandler interface {
	// Handles the new client connection
	Handle(connection *net.TCPConn)
}

// MessageEncoder interface to encode the out going messages
type MessageEncoder interface {
	Encode(data []byte) ([]byte, error)
}

// MessageDecoder interface to encode the incoming messages
type MessageDecoder interface {
	Decode(data []byte) ([]byte, error)
}
