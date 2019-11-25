package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// TCPServer to handle the incoming connection
type TCPServer struct {
	// Listen address
	ListenAddress *net.TCPAddr
	ConnHandler   ConnectionHandler
}

func (server *TCPServer) handleClient(connection *net.TCPConn) error {
	connection.SetKeepAlive(true)
	server.ConnHandler.Handle(connection)
	return nil
}

func interruptHandler() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	log.Println("CTRL-C; exiting")
	os.Exit(0)
}

// Start the TCP server
func (server *TCPServer) Start() error {
	tcpListener, err := net.ListenTCP("tcp", server.ListenAddress)
	if err != nil {
		log.Error(err)
		return err
	}
	defer tcpListener.Close()
	log.Printf("server started @ %s", tcpListener.Addr().String())

	go server.acceptClients(tcpListener)
	interruptHandler()
	return nil
}

// Accepts the new client connections and dispatches the handling
// to a new coroutine.
// TODO: Need to make this as queue based approach to define
// limits and get visibility.
func (server *TCPServer) acceptClients(tcpListener *net.TCPListener) {
	for {
		clientConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Print("error accepting client connection - ", err.Error())
			continue
		}

		log.Printf("new connection from %s", clientConn.RemoteAddr().String())

		go server.handleClient(clientConn)

	}
}
