package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a int32 = 128
	var b float32 = 12.1
	bits1 := math.Float32bits(b)
	fmt.Printf("%b\n", bits1)
	bits2 := strconv.FormatInt(int64(a), 2)
	fmt.Println(bits2)
	fmt.Printf("%b\n", a)
}
