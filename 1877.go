package main

import (
    "fmt"
)

func main() {
    a := 0
    b := 0

    fmt.Scan(&a)
    fmt.Scan(&b)

    if a%2 == 0 || b%2 == 1 {
        fmt.Println("yes")
    } else {
        fmt.Println("no")
    }
}
