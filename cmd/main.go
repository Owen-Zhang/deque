package main

import (
	"github.com/Owen-Zhang/deque"
)

func main() {
	qe := deque.NewDeque()

	for index := 0; index < 253; index++ {
		qe.PushBack(index)
	}
	qe.PushBack(255)
	qe.PushBack(256)
	qe.PushBack(258)

	qe.PushFront(260)
	qe.PopFront()
}
