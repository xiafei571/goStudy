package main

import (
	"fmt"

	"github.com/xiafei571/stringutil"
)

func main() {
	//第一个包 注意首字母大写
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Println()
	//冒泡排序
	fmt.Printf("%v", stringutil.BubbleSort([]int{14, 12, 3, 9, 4, 5, 8}))
}