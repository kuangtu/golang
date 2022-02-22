package main

import(
    "fmt"
)


func main() {
    u := [...]int{1,2,3,4}
    
    s := u[1:2]
    
    fmt.Printf("slice(len=%d,cap=%d):%v\n", len(s), cap(s), s)
    
    fmt.Println("array:", u)
        
    s = append(s, 5)
    fmt.Println("after append 5, array:", u)
    
    s = append(s, 6)
    fmt.Println("after append 6, array:", u)
    
    s = append(s, 7)
    fmt.Println("after append 7, array:", u)
    

}