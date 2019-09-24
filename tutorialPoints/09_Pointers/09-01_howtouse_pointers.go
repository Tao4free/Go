package main

import "fmt"

func main() {
    // actual variable declaration
    var a int = 20 
    // pointer variable declaration
    var ip *int

    // store address of a in pointer variable
    ip = &a

    fmt.Printf("Address of a variable: %x\n", &a)

    // address stored in pointer variable
    fmt.Printf("Address stored in ip variable: %x\n", ip)

    // access the value using the pointer
    fmt.Printf("Value of *ip variable: %d\n", *ip)
}
