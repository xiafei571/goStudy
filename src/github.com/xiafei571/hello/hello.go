package main

import (
	"fmt"

	"github.com/xiafei571/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println()
	fmt.Printf("%v", stringutil.BubbleSort([]int{14, 12, 3, 9, 4, 5, 8}))
}