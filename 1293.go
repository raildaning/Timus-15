package main

import (
    "fmt"
)

func main() {
    var a, b, n int
    fmt.Scan(&n)
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println((a * b) * 2 * n)
}
