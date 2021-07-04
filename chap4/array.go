package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(len(a))

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	//基本类型，数组比较:
	b1 := [4]byte{'a', 'b', 'c'}
	b2 := [4]byte{'a', 'b', 'c'}
	if b1 == b2 {
		fmt.Println("the two array equal")
	}
}
