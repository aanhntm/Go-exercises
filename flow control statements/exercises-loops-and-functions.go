package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = float64(1.0)
	var temp = float64(0.0)
	for i := 0; i < 10; i++ {
		temp = z
		z -= (z*z - x) / (2*z)
		if math.Sqrt(z) == math.Sqrt(temp){
			return(temp)
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(72))
}