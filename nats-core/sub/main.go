package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Welcome to the NATS Core subscriber!")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("only one single subject command line arg should be provided")
	}
	subject := args[0]

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("Subscriber %s is listening on %s\n", uuid.New(), subject)

	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		log.Printf("Received from %s:%s\n", m.Subject, string(m.Data))
	})
	if err != nil {
		log.Fatalf("error subscribing to %s:%v\n", subject, err)
	}

	// Exit the main goroutine but allow subscribe to continue running
	runtime.Goexit()
}
