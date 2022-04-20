# nats-demo
My NATS demo for the talk ["Using NATS for multi-cloud event streaming"](https://www.youtube.com/watch?v=AhnL5addsVo).

## Prerequisites

Please run these prerequisite steps
1. Go [installed](https://go.dev/doc/install)
1. Checkout code. 
1. Run `go mod download` in the terminal

## NATS Core demo
Run each command in a different terminal window or tab.
1. Start the NATS server `nats-server`. 
1. Start a publisher on a given subject eg. `go run nats-core/pub/main.go "payments.uk"`
1. Start a subscriber with a given subject eg. `go run nats-core/sub/main.go "payments.*"`

### Expected behaviour
The publisher will generate random payment messages and send them to the subscriber. As NATS core does not have any persistence, the messages published before the subscriber started up will be lost.

## JetStream demo
Run each command in a different terminal window or tab.
1. Start the JetStream server with the provided configuration file `nats-server –c jetstream/js.conf`. 
1. Start a publisher on a given subject eg. `go run jetstream/pub/main.go`
1. Start a push subscriber with a given subject eg. `go run jetstream/consumer-push/main.go "payments.*"`
1. Start a pull subscriber with a given subject eg. `go run jetstream/consumer-pull/main.go "payments.*"`

### Expected behaviour
The publisher will generate random payment messages and send them to the subscriber. As JetStream does have persistence, the messages published before the subscriber started up will be sent to the subscribers once they are started. The pull subscriber will read messages in batches of 3, while the push subscriber will be immediately caught up with messages 
