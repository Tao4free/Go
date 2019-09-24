package main

import "fmt"

func main() {
    var greeting = "Hellp world!"

    fmt.Printf("normal string: ")
    fmt.Printf("$s", greeting)
    fmt.Printf("\n")
    fmt.Printf("hex bytes: ")

    for i := 0; i < len(greeting); i++ {
        fmt.Printf("%x", greeting[i])
    }

    fmt.Print("\n")
    const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", sampleText)
    fmt.Printf("\n")
}
