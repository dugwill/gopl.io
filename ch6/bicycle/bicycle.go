// There are no examples in chap 6, so I made this to try a few thing

// Package bicycle
//!+bicycle
package bicycle

import (
	"math"
)

type Bicycle struct {
	Name        string
	Gears       int     // number of gears
	CurrentGear int     // current gear setting
	Cadence     float32 // pedel speed in revs per minute
	TireSize    int     // tire diameter in inches
}

//CalcSpeed take a point to a Bicycle and calculates the speed of the bike
//return in is MPH

func (b Bicycle) CalcSpeed() float32 {

	tireCircumference := 2 * math.Pi * float32(b.TireSize/2) //find circumference of current bycycle tire

	speedIPM := (float32(b.CurrentGear) * b.Cadence * tireCircumference) //calc speed in per min

	speedMPH := speedIPM * .00947

	return (speedMPH)

}
