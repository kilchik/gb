package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"fmt"
)

const (
	circuitStateOpen = "OPEN"
	circuitStateClosed = "CLOSED"
	circuitStateHalfOpen = "HALF-OPEN"
)

const (
	reqTimeout = 1 * time.Second
	tryCloseTimeout = 3 * time.Second
	triesBeforeSwitch = 3
)


func main() {
	orders := []string{
		"123432-234",
		"123432-234",
		"123432-234",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"254432-341",
		"254432-341",
		"254432-341",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"123432-234",
		"123432-234",
		"853432-332",
		"123432-234",
		"254432-341",
		"254432-341",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"123432-234",
		"123432-234",
		"853432-332",
		"123432-234",
		"254432-341",
		"254432-341",
		"853432-332",
		"853432-332",
		"853432-332",
		"254432-341",
		"123432-234",
		"123432-234",
		"853432-332",
		"123432-234",
	}

	ordersChan := make(chan string)

	proc := &Proc{circuitState: circuitStateClosed}
	go proc.ProcessPayments(ordersChan)
	for _, order := range orders {
		ordersChan <- order
		time.Sleep(1*time.Second)
	}
}

type Proc struct {
	circuitState string
	circuitOpenedAt time.Time
	failures, successes, total uint
}

func (p *Proc) process(order string) {
	p.total++

	switch p.circuitState {
	case circuitStateClosed:
		if err := pay(order); isTemp(err) {
			log.Printf(p.circuitState+": failed to pay for %q: %v", order, err)
			p.circuitState = circuitStateHalfOpen
			return
		}
		log.Printf(p.circuitState+": paid for %q", order)

	case circuitStateHalfOpen:
		if p.total % 2 != 0 {
			log.Printf(p.circuitState+": try order %q", order)
			if err := pay(order); isTemp(err) {
				log.Printf(p.circuitState + ": failed to pay for %q: %v", order, err)
				p.successes = 0
				p.failures++
				if p.failures >= triesBeforeSwitch {
					p.circuitState = circuitStateOpen
					p.failures = 0
					p.circuitOpenedAt = time.Now()
				}
			} else {
				p.failures = 0
				p.successes++
				if p.successes >= triesBeforeSwitch {
					p.circuitState = circuitStateClosed
				}
			}
			return
		}

		log.Printf(p.circuitState+": skip")

	case circuitStateOpen:
		if p.circuitOpenedAt.Add(tryCloseTimeout).Before(time.Now()) {
			log.Printf(p.circuitState + ": trying to close circuit...")
			if err := pay(order); isTemp(err) {
				log.Printf(p.circuitState + ": failed to pay for %q: %v", order, err)
				p.circuitOpenedAt = time.Now()
				return
			}

			p.successes = 1
			p.failures = 0
			p.circuitState = circuitStateHalfOpen
		}
	}
}

func (p *Proc) ProcessPayments(orders <-chan string)  {
	for order := range orders {
		p.process(order)
	}
}

var tempError = fmt.Errorf("temp error")

func isTemp(err error) bool {
	return err == tempError
}

func pay(order string) error {
	cli := http.Client{Timeout: reqTimeout}
	buf := bytes.NewBuffer([]byte(order))
	resp, err := cli.Post("http://localhost:5301/pay", "", buf)
	if err != nil {
		log.Printf("pay: %v", err)
		return tempError
	}
	if resp.StatusCode / 100 == 5 {
		return tempError
	}
	return nil
}
