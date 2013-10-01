package main

import (
    "fmt"
)

var stackList map[int][]int

func push(stackId int, object int) {
    stack, ok := stackList[stackId]
    if !ok {
        stack = make([]int, 0)
    }
    stackList[stackId] = append(stack, object)
}

func pop(stackId int) int {
    stack := stackList[stackId]
    object, stack := stack[len(stack)-1], stack[:len(stack)-1]
    stackList[stackId] = stack
    return object
}

func main() {
    var numb int
    var stackId int
    var object int
    var op string
    stackList = make(map[int][]int)
    fmt.Scan(&numb)
    for i := 1; i <= numb; i++ {
        fmt.Scan(&op)
        if op == "PUSH" {
            fmt.Scan(&stackId)
            fmt.Scan(&object)
            push(stackId, object)
        } else if op == "POP" {
            fmt.Scan(&stackId)
            object = pop(stackId)
            fmt.Println(object)
        }
    }
}
