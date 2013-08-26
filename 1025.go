package main

import (
	"fmt"
	"sort"
)

func main() {
	nbGroups := 0
	groups := make([]int, 0)
	sum := 0
	fmt.Scan(&nbGroups)
	for i := 0; i < nbGroups; i++ {
		var i int
		fmt.Scan(&i)
		groups = append(groups, i)
	}
	sort.Ints(groups)
	lowerGroups := groups[:(len(groups)/2 + 1)]
	for _, value := range lowerGroups {
		if value == 1 {
			sum += 1
			continue
		}
		sum += value/2 + 1
	}
	fmt.Println(sum)
}
