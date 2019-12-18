package main

import (
	"fmt"
	"math"
)

func main() {
	rate := 109.90
	cost := 480.81
	cost *= rate
	fmt.Println(math.Round(cost))
}
