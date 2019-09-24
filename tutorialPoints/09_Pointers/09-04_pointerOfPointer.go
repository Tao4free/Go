package main 

import "fmt"

func main() {
    var a int
    var ptr *int
    var pptr **int

    a = 3000

    // take the address of var
    ptr = &a

    // takke the address of ptr using address of operator &
    pptr = &ptr


    // take the value using pptr
    fmt.Printf("value of a = %d\n", a)
    fmt.Printf("value available at *ptr = %d\n", *ptr)
    fmt.Printf("value available at **pptr = %d\n", **pptr)

    **pptr /= 2
    fmt.Printf("After division: %d\n", a)
}
