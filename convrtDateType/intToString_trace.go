package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.Itoa(12)
	s := strconv.FormatInt(int64(33333), 10)
	fmt.Printf("%q %q\n", a, s)
}
