package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fp, err := os.OpenFile("hellworld.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()

	fpread := bufio.NewReaderSize(fp, 1024)

	lines, b, err := fpread.ReadLine()

	if err != nil {
		fmt.Println("fp read line error")
		return
	}

	if b == true {
		fmt.Println("lines too large")

		return
	}

	fmt.Println("read line is:", string(lines))
}
