package main

import(
    "fmt"
    "unsafe"
)


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
    
    var arr = [5]int{1, 2, 3, 4, 5}
    fmt.Println("数组长度:", len(arr))
    fmt.Println("数组大小:", unsafe.Sizeof(arr))
}
