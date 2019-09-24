package main

import "fmt"

func main() {
    var c1, c2, c3 chan int
    var i1, i2 int
    fmt.Printf("c1 = %s\n", c1)
    fmt.Printf("c2 = %s\n", c2)
    fmt.Printf("c3 = %s\n", c3)

    select {
        case i1 = <- c1:
            fmt.Printf("received", i1, "from c1\n")
        case c2 <- i2:
            fmt.Printf("sent", i2, "to c2\n")
        case i3, ok := (<-c3):
            if ok {
                fmt.Print("Received", i3, "from c3\n")
            } else {
                fmt.Print("c3 is closed\n")
            }
        default:
            fmt.Printf("no communication\n")
    }
}
