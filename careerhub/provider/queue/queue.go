package queue

import (
	"fmt"
)

type Queue interface {
	Send(message *string) error
}

type FakeQueue struct {
}

func (fq *FakeQueue) Send(message *string) error {

	fmt.Printf("%s\n\n", *message)
	return nil
}
