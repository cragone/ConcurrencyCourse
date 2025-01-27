package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed a random number generator
	rand.NewSource(time.Now().UnixNano())

	//print a welcome message
	color.Yellow("Welcome too the Sleeping barber, Barbershop")
	color.Yellow("-------------------------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create the barber shop
	shop := BarberShop{ShopCapacity: seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Lorena")

	//start the barber shop as a go routine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	//add clients

	//block until the barbershop is closed
	time.Sleep(5 * time.Second)

}
