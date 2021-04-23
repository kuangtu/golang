package main

import(
    "os"
    "fmt"
    "encoding/json"
)

type Person struct {
    Name Name
    Email [] Email
}

type Name struct {
    Family string
    Personal string
}

type Email struct {
    Kind string
    Address string
}


func main() {
    //结构体初始化
    person := Person {
            Name: Name{Family: "newmarch", Personal: "Jan"},
            Email: []Email{Email{Kind: "home", Address: "Jan@newmarch.name"},
                        Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}
                        }
                    }
    saveJSON("person.json", person)
    
}


func saveJSON(fileName string, key interface{}) {
    outFile, err := os.Create(fileName)
    checkError(err)
    encoder := json.NewEncoder(outFile)
    err := encoder.Encode(key)
    checkError(err)
    outFile.Close()
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
