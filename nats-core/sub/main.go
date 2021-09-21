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
	log.Printf("subscriber %s is listening on %s\n", uuid.New(), subject)

	var scount int
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		scount++
		log.Printf("[%d] received from %s: %s\n", scount, m.Subject, string(m.Data))
	})
	if err != nil {
		log.Fatalf("error subscribing to %s:%v\n", subject, err)
	}

	runtime.Goexit()
}
