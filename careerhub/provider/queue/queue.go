package queue

import (
	"fmt"
	"io"
)

type Queue interface {
	Send(rc io.Reader) error
}

type FakeQueue struct {
}

func (fq *FakeQueue) Send(rc io.Reader) error {
	b, _ := io.ReadAll(rc)
	// fmt.Print("Done sending match data\n")
	fmt.Printf("%s\n\n", string(b))
	return nil
}
