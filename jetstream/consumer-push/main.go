package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Welcome to the NATS JetStream consumer!")
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
	log.Printf("push Consumer %s is listening on %s\n", uuid.New(), subject)

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating jetstream Context: ", err)
	}

	var scount int
	js.Subscribe(subject, func(m *nats.Msg) {
		scount++
		log.Printf("[%d] received from %s: %s\n", scount, m.Subject, string(m.Data))
	})

	runtime.Goexit()
}
