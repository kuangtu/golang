package main

import (
    "fmt"
    "io"
)


type MyInt *int
func (r MyInt) String() string { 
    return fmt.Sprintf("%d", *(*int)(r))
}

type MyReader io.Reader
func (r MyReader) Read(p []byte) (int, error) {
    return r.Read(p)
}

func main() {
    values := []int{1, 2, 3, 4}
    fmt.Println(sum(values...)) // "10"
}