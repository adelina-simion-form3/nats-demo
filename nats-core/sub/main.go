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
		log.Fatal("Only one single subject command line arg should be provided")
	}
	subject := args[0]
	
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Error connecting to NATS Server", err)
	}
	fmt.Printf("Subscriber %s is listening on %s\n", uuid.New(), subject)

	var scount int
	nc.Subscribe(subject, func(m *nats.Msg) {
		scount++
		fmt.Printf("[%d] Received from %s: %s\n", scount, m.Subject, string(m.Data))
	})

	runtime.Goexit()
}
