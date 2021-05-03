package main

import (
    "os"
    "log"
    "fmt"
)

func main() {
    filepath := "test.txt"

    newFile, err := os.Create(filepath)
    if err != nil {
        log.Fatal(err)
    }
    //write string
    newFile.Write([]byte("hello world"))
    newFile.Close()
    
    //open and read file
    bytebuffer := make([]byte, 16)   
    newFile2, err := os.Open(filepath)
    nbytes, err:= newFile2.Read(bytebuffer)
    fmt.Printf("read len:%d str is:%s", nbytes, string(bytebuffer))
    defer newFile2.Close()
}