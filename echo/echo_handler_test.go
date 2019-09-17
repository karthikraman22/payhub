package echo_test

import (
	"net"
	"os"
	"testing"

	"achuala.in/payhub/server"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestEchoServer(*testing.T) {
	listenAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort("127.0.0.1", "18081"))
	if err != nil {
		log.Panic(err)
	}
	echoServer := &server.TCPServer{ListenAddress: listenAddr, ConnHandler: &server.EchoHandler{}}

	err = echoServer.Start()
	if err != nil {
		log.Error(err)
	}
}
