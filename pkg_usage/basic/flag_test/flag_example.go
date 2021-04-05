package main

import (
    "flag"
    "fmt"
)

func main() {
    wordPtr := flag.String("word", "foo", "a string")
    numberPtr := flag.Int("numb", 10, "an int")
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    
    flag.Parse()
    
    fmt.Println("word", *wordPtr)
    fmt.Println("numb", *numberPtr)
    fmt.Println("svar", svar)
    
}