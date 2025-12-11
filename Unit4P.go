// Author: William Provost
// Version: 1.0.0
// Date: 2025-12-10
// Fileoverview: This program keeps track of car stats and simulates drive/fill/oil change

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initial data
var makeCar string = "UsedCarBrand"
var model string = "ModelX"
var carColor string = "Silver"
var odometer int = 65000    // starting mileage
var oilChangeKM int = 65000 // last oil change
var gasCost [10]float64
var fillIndex int = 0

// carStats returns a string of car info
func carStats() string {
	return fmt.Sprintf(
		"Car Make: %s\nCar Model: %s\nCar Colour: %s\nOdometer: %d km\nLast oil change at: %d km\nNumber of recorded fill-ups: %d",
		makeCar, model, carColor, odometer, oilChangeKM, fillIndex)
}

// wrapCar prompts user for new colour
func wrapCar() string {
	var newColour string
	fmt.Print("Enter a new colour to wrap the car with: ")
	fmt.Scanln(&newColour)
	if newColour == "" {
		return carColor
	}
	return newColour
}

// drive generates random km (100..1000) and updates odometer
func drive() int {
	rand.Seed(time.Now().UnixNano())
	driven := rand.Intn(901) + 100 // 100..1000
	odometer += driven
	return driven
}

// fillUp prompts for cost and stores in gasCost array
func fillUp() {
	if fillIndex >= len(gasCost) {
		fmt.Println("Gas cost array is full; cannot record more fill-ups.")
		return
	}
	var cost float64
	fmt.Print("Enter how much you paid for the gas to fill up: ")
	fmt.Scanln(&cost)
	gasCost[fillIndex] = cost
	fmt.Printf("Recorded fill-up of $%.2f at slot %d.\n", cost, fillIndex)
	fillIndex++
}

// displayCostToFillUp shows all recorded costs and returns average
func displayCostToFillUp() float64 {
	if fillIndex == 0 {
		fmt.Println("No fill-up records.")
		return 0
	}
	fmt.Println("\nRecorded fill-up costs:")
	sum := 0.0
	for i := 0; i < fillIndex; i++ {
		fmt.Printf("Fill-up %d: $%.2f\n", i+1, gasCost[i])
		sum += gasCost[i]
	}
	avg := sum / float64(fillIndex)
	fmt.Printf("Average cost per fill-up: $%.2f\n", avg)
	fmt.Printf("\n(Info) Returned average cost: $%.2f\n", avg)
	return avg
}

// oilChange checks if vehicle needs oil change and updates oilChangeKM
func oilChange() bool {
	if odometer-oilChangeKM >= 5000 {
		oilChangeKM = odometer
		fmt.Println("\nAn oil change was done.")
		return true
	}
	fmt.Println("\nYour car does not need an oil change.")
	return false
}

func main() {
	// initial fill-up 
	gasCost[fillIndex] = 74.0
	fillIndex++

	fmt.Println("Initial car stats:")
	fmt.Println(carStats())

	// wrap the car
	carColor = wrapCar()
	fmt.Printf("\nCar colour updated to: %s\n", carColor)

	// drive
	drivenKm := drive()
	fmt.Printf("\nYou drove %d km. New odometer: %d km.\n", drivenKm, odometer)

	// fill up
	fillUp()

	// display fill-up costs and average
	_ = displayCostToFillUp()

	// check oil change
	oilChange()

	fmt.Println("\nFinal car stats:")
	fmt.Println(carStats()) // <-- this shows odometer, oilChangeKM, and fillIndex

	fmt.Println("\nDone.")
}
