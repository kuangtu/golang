package main

import "fmt"



func sum(values... int) int {
    total := 0
    for _, val := range values {
    total += val
    }
    
    return total
}

func main() {
    values := []int{1, 2, 3, 4}
    fmt.Println(sum(values...)) // "10"
}