package main

import (
	"fmt"
)

func main() {

	var a = [...]int{1, 2, 3, 4}
	fmt.Println(a[1])

	var b = []int{5, 6, 7, 8}
	s := a[:]
	s = b
	fmt.Printf("%v", s)

	var header [4]byte
	slice1 := header[:]
	//对于切片的修改，会反映到底层数组中
	copy(slice1, []byte("test"))
	fmt.Println(header)

}
