package echo

import (
	"bufio"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

// EchoHandler which responds with the same incoming message
type EchoHandler struct {
}

// Handle implements the echo message
func (eh *EchoHandler) Handle(connection *net.TCPConn) {
	defer connection.Close()
	buf := bufio.NewReader(connection)
	for {
		msg, err := buf.ReadString('\n')
		if err == io.EOF {
			// Client connection closed
			break
		}
		if err != nil {
			log.Errorf("error while reading data from client %s,  error code %v", connection.RemoteAddr().String(), err)
			break
		}
		_, err = connection.Write([]byte(msg))
		if err != nil {
			log.Errorf("error while sending data to client %s,  error code %v", connection.RemoteAddr().String(), err)
			break
		}
	}
	log.Printf("closing connection %s", connection.RemoteAddr().String())
}
