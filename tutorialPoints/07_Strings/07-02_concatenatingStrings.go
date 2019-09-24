package main

import (
    "fmt"
    "strings"
)

func main() {
    greetings := []string{"hello", "world!"}
    fmt.Println(strings.Join(greetings, " "))
}
