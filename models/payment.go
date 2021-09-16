package models

import (
	"fmt"
	"math/rand"
)

const maxRandomAmount int = 500
const currency string = "GBP"

type Payment struct {
	amount   int
	currency string
	sender   string
	receiver string
}

func (p Payment) String() string {
	return fmt.Sprintf("%s sent a payment of %d %s to %s", p.sender, p.amount, p.currency, p.receiver)
}

func GetRandomPayment() string {
	a := rand.Intn(maxRandomAmount)
	s := rand.Intn(maxRandomAmount)
	r := rand.Intn(maxRandomAmount)
	p := Payment{a, currency, fmt.Sprintf("SENDER%d", s), fmt.Sprintf("RECEIVER%d", r)}
	return p.String()
}
