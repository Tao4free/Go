// https://golang.org/pkg/math/

package main

import (
	"fmt"
	"math"
)

func main() {
	p := math.Round(10.4)
	fmt.Printf("%.1f\n", p)

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n)
}
