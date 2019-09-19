package client_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"achuala.in/payhub/iso8583/client"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"

	lib8583 "github.com/mofax/iso8583"
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestIso8583Client(*testing.T) {
	client, err := client.Connect("127.0.0.1", "18081", "../spec1987.yml")
	if err != nil {
		log.Error(err)
	}
	rand.Seed(100)
	for i := 1; i < 100000; i++ {
		isoMsg := lib8583.NewISOStruct("../spec1987.yml", true)
		isoMsg.AddMTI("0200")
		isoMsg.AddField(2, fmt.Sprintf("%016d", rand.Int63n(1e16)))
		err = client.Send(&isoMsg)
	}
	time.Sleep(100 * time.Second)
}
