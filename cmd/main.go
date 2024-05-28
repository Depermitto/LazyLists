package main

import (
	"fmt"
	"github.com/Depermitto/LazyLists/lazy"
)

func main() {
	fruits := []string{"banana", "apple", "orange", "jam", "pineapple"}
	for i, fruit := range lazy.Enumerate[string](fruits).Skip(2).StepBy(2).Collect() {
		fmt.Println(i, fruit)
	}
}
