package main

import (
	"flag"
	"log"
	"time"

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
		log.Fatal("Only one single subject command line arg should be provided")
	}
	subject := args[0]

	nc, err := nats.Connect(jetstreamURL)
	if err != nil {
		log.Fatal("error connecting to nats server:", err)
	}
	log.Printf("pull consumer %s is listening on %s\n", uuid.New(), subject)

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating jetstream context", err)
	}

	// MONITOR consumer without any acknowledgement; see docs for other options
	sub, err := js.PullSubscribe(subject, "MONITOR")
	if err != nil {
		log.Fatal("error creating pull subscriber", err)
	}

	for {
		msgs, err := sub.Fetch(3)
		if err != nil {
			log.Fatal("error fetching message from pull subscriber", err)
		}
		for _, m := range msgs {
			log.Printf("Received from %s: %s\n", m.Subject, string(m.Data))
		}
		// Poll every 5 second
		time.Sleep(5 * time.Second)
	}
}
