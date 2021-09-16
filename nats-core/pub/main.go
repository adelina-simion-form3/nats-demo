package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-demo/models"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Welcome to the NATS Core publisher!")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	subject := args[0]

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Error connecting to NATS Server", err)
	}
	fmt.Printf("Publisher %s is publishing on %s\n", uuid.New(), subject)
	

	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	var pcount int
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				pcount++
				p := models.GetRandomPayment()
				fmt.Printf("[%d] Publishing on %s:%s\n", pcount, subject, p)
				nc.Publish(subject, []byte(p))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()
}
