package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//dak
	fp, err := os.OpenFile("hellworld.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)

		return
	}

	fpwrite := bufio.NewWriterSize(fp, 1024)
	fpwrite.Write([]byte("hello world\n"))
	fmt.Println(fpwrite.Available())
	fmt.Println(fpwrite.Size())
	fmt.Println(fpwrite.Buffered())
	fpwrite.Flush()
}
