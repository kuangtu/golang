package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("create file demo")
	name := "srctest.txt"
	fp, err := os.Create(name)
	if err != nil {
		fmt.Println("create file error:", err)
	}
	fp.Close()

	src := "srctest.txt"
	dst := "dsttest.txt"
	renerr := os.Rename(src, dst)
	if renerr != nil {
		fmt.Println("rename error:", err)
	}
}
