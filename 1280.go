package main

import (
    "fmt"
)

var depList map[int][]int
var processedSubjects map[int]nothing

type nothing struct{}

func addDep(subject int, dep int) {
    depList[subject] = append(depList[subject], dep)
}

func processSubjectsList(subjectList []int) bool {
    processedSubjects = make(map[int]nothing)

    for _, subject := range subjectList {
        if depList[subject] == nil {
            processedSubjects[subject] = nothing{}
        } else {
            for _, dep := range depList[subject] {
                processedSubjects[subject] = nothing{}
                _, ok := processedSubjects[dep]

                if !ok {
                    return false
                }
            }
        }
    }
    return true
}

func main() {
    var n int
    var m int
    depList = make(map[int][]int)
    fmt.Scan(&n)
    fmt.Scan(&m)

    for i := 1; i <= m; i++ {
        var dep int
        var subject int
        fmt.Scan(&dep)
        fmt.Scan(&subject)
        addDep(subject, dep)
    }

    subjectList := make([]int, 0)

    for i := 1; i <= n; i++ {
        var a int
        fmt.Scan(&a)
        subjectList = append(subjectList, a)
    }

    if processSubjectsList(subjectList) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
