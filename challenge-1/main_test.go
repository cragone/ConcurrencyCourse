package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("pie", &wg)
	wg.Wait()

	if msg != "pie" {
		t.Errorf("output was not pie")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "beta"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "beta") {
		t.Errorf("expected beta but it was not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("expected hello universe not there")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("expected hello cosmos not there")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("expected hello world not there")
	}
}
