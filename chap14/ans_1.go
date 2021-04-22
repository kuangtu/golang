package main

import(
    "encoding/asn1"
    "os"
    "fmt"
)

func main() {

    mdata, err := ans1.Marshal(13)
    
    var n int
    
    _, err1 := asn1.UnMarshal(mdata, &n)

    checkError(err1)
    
    fmt.Println("After marshal/unmarshal: ", n)
}


func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
