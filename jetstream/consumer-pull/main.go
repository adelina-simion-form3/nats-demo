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
		log.Fatal("Only one single subject command line arg should be provided")
	}
	subject := args[0]

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS server", err)
	}
	log.Printf("pull Consumer %s is listening on %s\n", uuid.New(), subject)

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating jetstream context", err)
	}

	var scount int
	// Ask around what these strings are defined as - should be regular or durable
	sub, err := js.PullSubscribe(subject, "MONITOR")
	if err != nil {
		log.Fatal("error creating pull subscriber", err)
	}
	// Cleanup
	defer func() {
		// Unsubscribe
		sub.Unsubscribe()
		// Drain
		sub.Drain()
	}()

	msgs, err := sub.Fetch(3)
	if err != nil {
		log.Fatal("error fetching message from pull subscriber", err)
	}
	for _, m := range msgs {
		scount++
		log.Printf("[%d] received from %s: %s\n", scount, m.Subject, string(m.Data))
	}

	runtime.Goexit()
}
