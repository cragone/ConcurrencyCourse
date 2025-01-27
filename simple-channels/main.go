package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			//do something.
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	//create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit!)")

	for {
		fmt.Println("-> ")

		var userInput string

		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		//wait for a response.

		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done, closing channels.")
	close(ping)
	close(pong)
}
