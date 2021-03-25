package main

import (
        _ "bytes"
        _ "net/http"
        _ "time"
        "fmt"
        "github.com/golang/groupcache"
)

func main() {
    fmt.Println("Hello world.")
    me := "http://10.0.0.1"
    peers := groupcache.NewHTTPPool(me)
    peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")
    

    
}