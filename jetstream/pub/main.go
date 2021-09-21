package main

import (
	"flag"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-demo/models"
	"github.com/nats-io/nats.go"
)

const maxMessages = 500

func main() {
	log.Println("Welcome to the NATS JetStream publisher!")
	flag.Parse()
	args := flag.Args()
	log.Println(args)
	subject := args[0]

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("publisher %s is publishing on %s\n", uuid.New(), subject)

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating JetStream Context", err)
	}

	// Simple Stream Publisher
	for i := 0; i < maxMessages; i++ {
		p := models.GetRandomPayment()
		log.Printf("[%d] publishing on %s:%s\n", i, subject, p)

		js.Publish(subject, []byte(p))
		time.Sleep(2 * time.Second)
	}
}
