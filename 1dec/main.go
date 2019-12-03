package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Fuel-calculator")
	scanner := bufio.NewScanner(os.Stdin)

	var totalFuelRequirement float64

	for scanner.Scan() {
		moduleMass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return
		}

		totalFuelRequirement += calcFuelRecur(moduleMass)
		fmt.Println(totalFuelRequirement)
	}
	fmt.Println(totalFuelRequirement)
}

func calcFuelRecur(moduleMass float64) (fuel float64) {
	currentFuel := (math.Floor(moduleMass / 3)) - 2
	if currentFuel <= 0 {
		return 0
	}
	return currentFuel + calcFuelRecur(currentFuel)
}
