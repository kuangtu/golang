package main

import "fmt"

func main() {
    var a [3]int
    a = append(a, 1)
    a = append(a, 2)
    a = append(a, 3, 4)
    fmt.Println(a)
}