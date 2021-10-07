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
const subject = "payments.uk"

// FOR DEMO to not clash on default port
const jetstreamURL = "nats://127.0.0.1:5222"

func main() {
	log.Println("Welcome to the NATS JetStream publisher!")

	// Connect to a server
	nc, err := nats.Connect(jetstreamURL)
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
		if err := js.DeleteStream(streamName); err != nil {
			log.Fatal("error deleting stream:", err)
		}
	}()

	log.Printf("publisher %s is publishing on %s\n", uuid.New(), subject)

	// Simple Stream Publisher
	for i := 0; i < maxMessages; i++ {
		p := models.GetRandomPayment(i)
		log.Printf("Publishing on %s:%s\n", subject, p)

		if _, err := js.Publish(subject, []byte(p)); err != nil {
			log.Fatal("error publishing:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
