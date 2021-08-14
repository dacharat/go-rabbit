package rabbit

import (
	"github.com/dacharat/go-rabbit/cmd/server/queue"
)

func Publish() error {
	return queue.Publish("hello", []byte("Hello World!!"))
}
