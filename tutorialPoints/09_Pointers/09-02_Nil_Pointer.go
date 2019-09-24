package main

import "fmt"

func main() {
    var ptr *int

    fmt.Printf("The value of ptr is %x\n", ptr)


    if (ptr == nil) {
        fmt.Printf("This pointer is Nil Pointer\n")
    }
}
