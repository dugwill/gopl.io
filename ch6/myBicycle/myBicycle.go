package main

import (
	"fmt"

	"github.com/dugwill/gopl.io/ch6/bicycle"
)

func main() {

	var speed float32

	singleSpeed:= new(bicycle.Bicycle)
	singleSpeed.Name = "motocross"
	singleSpeed.Gears = 1
	singleSpeed.CurrentGear = 1
	singleSpeed.TireSize = 14
	singleSpeed.Cadence = 60

	speed = singleSpeed.CalcSpeed()

	fmt.Printf("%v with Tiresize: %v in. has a tire circumfrence of %v in.\n", singleSpeed.Name,singleSpeed.TireSize, singleSpeed.CalcCircumference())

	fmt.Printf("%v with Cadence %v: Speed is: %v MPH\n", singleSpeed.Name, singleSpeed.Cadence, speed)

	fmt.Printf("%v\n",*singleSpeed)

	tenSpeed:= new(bicycle.Bicycle)
	tenSpeed.Name = "roadbike"
	tenSpeed.Gears = 10
	tenSpeed.CurrentGear = 5
	tenSpeed.TireSize = 24
	tenSpeed.Cadence = 40

	speed = tenSpeed.CalcSpeed()

	fmt.Printf("%v with Tiresize: %v in. has a tire circumfrence of %v in.\n", tenSpeed.Name,tenSpeed.TireSize, tenSpeed.CalcCircumference())

	fmt.Printf("%v with Cadence %v: Speed is: %v MPH\n", tenSpeed.Name, tenSpeed.Cadence, speed)

	fmt.Printf("%v\n",*tenSpeed)

}
