package main
import "fmt"


func init() {
    fmt.Println("invoke init")
}

func main() { 
    init()
}