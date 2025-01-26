package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	var name string
	var tone string
	color.Cyan("What is your name?")
	fmt.Scanln(&name)
	color.Cyan("cool %s", name)
	time.Sleep(time.Second * 5)
	color.Cyan("whats your favorite color?")
	fmt.Scanln(&tone)
	color.Blue("Nice, I love %s too!", tone)
}
