package main

import "fmt"

func add(x,y int) int {
    return x  + y
}

func swap(x,y *int){
    var c int
    c = *x
    *x = *y
    *y = c

}

func add(a,b int64) (c ,d int64) {
    x = a + b
    y = a - b
    c = x * 2
    d = y * 2
}

func main() {
    var a = 1
    var b = 2
    var c = add(a,b)
    fmt.Printf("%d\n", c)
}