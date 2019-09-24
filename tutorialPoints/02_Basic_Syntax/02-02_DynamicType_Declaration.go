package main

import "fmt"

func main() {
    var x float64 = 20.0

    y := 42
    fmt.Println(x)
    fmt.Println(y)
    fmt.Printf("%f is of type %T\n", x, x)
    fmt.Printf("y is of type %T\n", y)
}
