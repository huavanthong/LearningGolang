package main

import (
	"fmt"
	"time"
)

type CircuitBreaker struct {
	name                string
	maxFailures         int
	timeout             time.Duration
	state               string
	consecutiveFailures int
	lastFailureTime     time.Time
}

func (cb *CircuitBreaker) checkState() {
	if cb.consecutiveFailures >= cb.maxFailures {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = "half-open"
		} else {
			cb.state = "open"
		}
	} else {
		cb.state = "closed"
	}
}

func (cb *CircuitBreaker) Execute(request func() error) error {
	if cb.state == "open" {
		if time.Since(cb.lastFailureTime) > cb.timeout {
			cb.state = "half-open"
		} else {
			return fmt.Errorf("Circuit breaker %s is open", cb.name)
		}
	}

	err := request()
	if err != nil {
		cb.consecutiveFailures++
		cb.lastFailureTime = time.Now()
		cb.checkState()
		return err
	}

	cb.consecutiveFailures = 0
	cb.checkState()
	return nil
}

func main() {

	cb := &CircuitBreaker{
		name:        "serviceA",
		maxFailures: 3,
		timeout:     time.Second * 10,
	}

	err := cb.Execute(func() error {
		// Kết nối đến hệ thống A và thực hiện một yêu cầu
	})
	if err != nil {
		// Xử lý lỗi
	}

}
