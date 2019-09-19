package server_test

import (
	"os"
	"testing"

	"achuala.in/payhub/iso8583/server"

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

func TestIso8583Server(*testing.T) {
	isoServer := &server.IsoServer{Host: "127.0.0.1", Port: "18081", SpecFile: "../spec1987.yml"}

	err := isoServer.Start()
	if err != nil {
		log.Error(err)
	}
}
