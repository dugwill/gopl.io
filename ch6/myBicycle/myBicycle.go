package main

import (
	"fmt"

	"github.com/dugwill/gopl.io/ch6/bicycle"
)

func main() {
	var singleSpeed bicycle.Bicycle
	singleSpeed.Name = "motocross"
	singleSpeed.Gears = 1
	singleSpeed.CurrentGear = 1
	singleSpeed.TireSize = 14
	singleSpeed.Cadence = 60

	speed := singleSpeed.CalcSpeed()

	fmt.Printf("%v with Cadence %v: Speed is: %v\n", singleSpeed.Name, singleSpeed.Cadence, speed)

}
