package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, word!"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("incorrect message")
	}
}
