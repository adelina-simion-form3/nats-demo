package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// FOR DEMO to not clash on default port
const jetstreamURL = "nats://127.0.0.1:5222"

func main() {
	log.Println("Welcome to the NATS JetStream consumer!")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("only one single subject command line arg should be provided")
	}
	subject := args[0]

	nc, err := nats.Connect(jetstreamURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("push consumer %s is listening on %s\n", uuid.New(), subject)

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating jetstream Context: ", err)
	}

	js.Subscribe(subject, func(m *nats.Msg) {
		log.Printf("Received from %s: %s\n", m.Subject, string(m.Data))
	})
	
	// Exit the main goroutine but allow subscribe to continue running
	runtime.Goexit()
}
