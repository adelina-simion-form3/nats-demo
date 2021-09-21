package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nats-demo/models"
	"github.com/nats-io/nats.go"
)

const maxMessages = 500
const subject = "payments.uk"

func main() {
	log.Println("Welcome to the NATS Core publisher!")

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error connecting to NATS Server", err)
	}
	log.Printf("publisher %s is publishing on %s\n", uuid.New(), subject)

	// Publish up to max messages every 2 seconds
	for i := 0; i < maxMessages; i++ {
		p := models.GetRandomPayment()
		log.Printf("[%d] publishing on %s:%s\n", i, subject, p)

		if err := nc.Publish(subject, []byte(p)); err != nil {
			log.Fatal("error publishing:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
