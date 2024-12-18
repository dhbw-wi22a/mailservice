package workers

import (
	"fmt"
)

type Queue struct {
	channel chan EmailRequest
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		channel: make(chan EmailRequest, capacity),
	}
}

func (q *Queue) Add(email EmailRequest) {
	select {
	case q.channel <- email: // Add email to the channel
		fmt.Printf("Email added to queue: %s\n", email.Recipient)
	default:
		fmt.Println("Queue is full. Dropping email request.")
	}
}

func (q *Queue) Get() (EmailRequest, bool) {
	select {
	case email := <-q.channel:
		return email, true
	default:
		return EmailRequest{}, false
	}
}

func (q *Queue) Length() int {
	return len(q.channel)
}
