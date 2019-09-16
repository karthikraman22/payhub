package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	rtformatter "github.com/banzaicloud/logrus-runtime-formatter"
	nats "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/vys/go-humanize"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func init() {
	formatter := rtformatter.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
func interruptHandler() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	log.Println("CTRL-C; exiting")
	os.Exit(0)
}

func goRuntimeStats() {
	m := &runtime.MemStats{}
	log.Print("Hey")
	for {
		time.Sleep(2 * time.Second)
		log.Info("# goroutines: ", runtime.NumGoroutine())
		runtime.ReadMemStats(m)
		log.Info("Memory Acquired: ", humanize.Bytes(m.Sys))
		log.Info("Memory Used    : ", humanize.Bytes(m.Alloc))
		log.Info("# malloc       : ", m.Mallocs)
		log.Info("# free         : ", m.Frees)
		log.Info("GC enabled     : ", m.EnableGC)
		log.Info("# GC           : ", m.NumGC)
		log.Info("Last GC time   : ", m.LastGC)
		log.Info("Next GC        : ", humanize.Bytes(m.NextGC))
		//runtime.GC()
	}
}

func main() {
	go goRuntimeStats()

	servers := "nats://localhost:4222, nats://localhost:5222, nats://localhost:6222"

	// Connect to a server
	nc, err := nats.Connect(servers)
	if err != nil {
		log.WithFields(log.Fields{"servers": servers}).Error(err)
	}
	// Simple Publisher
	err = nc.Publish("foo", []byte("Hello World"))
	if err != nil {
		fmt.Print(err)
	}
	// Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Responding to a request message
	nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// Simple Sync Subscriber
	sub, _ := nc.SubscribeSync("foo")
	m, _ := sub.NextMsg(5000)
	fmt.Println(m)
	// Channel Subscriber
	ch := make(chan *nats.Msg, 64)
	sub, _ = nc.ChanSubscribe("foo", ch)
	msg := <-ch

	fmt.Print(msg)
	// Unsubscribe
	sub.Unsubscribe()

	// Drain
	sub.Drain()

	// Requests
	msg, _ = nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// Replies
	nc.Subscribe("help", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("I can help!"))
	})

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()

}
