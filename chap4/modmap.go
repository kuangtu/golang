package main

import(
    "fmt"
)

func modMapValue(m map[int]int) {
    
    //修改map
    m[2] = 15
   
}


func main() {
   
    //init
    m := map[int]int{
    1: 11,
    2: 12,
    3: 13,
    }

    modMapValue(m)
    
    v, ok := m[2]
    if ok {
        fmt.Println(v)
    }
}