package main

import "fmt"

func main() {
    var a int = 100
    var b int = 200

    fmt.Printf("Before swap, value of a: %d\n", a)
    fmt.Printf("Before swap, value of b: %d\n", b)

    swap(&a, &b)

    fmt.Printf("After swap, value of a: %d\n", a)
    fmt.Printf("After swap, value of b: %d\n", b)
}

func swap(x *int, y *int) {
    var temp int

    // save the value at address
    temp = *x 
    // put y into x
    *x = *y
    // put temp into y
    *y = temp
}
