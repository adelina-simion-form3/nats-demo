package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-demo/models"
	"github.com/nats-io/nats.go"
)

const maxMessages = 500
const streamName = "payments"

func main() {
	log.Println("Welcome to the NATS JetStream publisher!")
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}

	// Create JetStream Context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error creating jetstream context:", err)
	}
	
	// Create a Stream
	if _, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{fmt.Sprintf("%s.*", streamName)},
	}); err != nil {
		log.Fatal("error adding stream:", err)
	}

	//Cleanup
	defer func() {
		// Delete Stream
		js.DeleteStream(streamName)
	}()

	subject := "payments.uk"
	log.Printf("publisher %s is publishing on %s\n", uuid.New(), subject)

	// Simple Stream Publisher
	for i := 0; i < maxMessages; i++ {
		p := models.GetRandomPayment()
		log.Printf("[%d] publishing on %s:%s\n", i, subject, p)

		js.Publish(subject, []byte(p))
		time.Sleep(2 * time.Second)
	}
}
