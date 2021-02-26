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

func main() {
    var a = 1
    var b = 2
    var c = add(a,b)
    fmt.Printf("%d\n", c)
}