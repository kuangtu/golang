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

//绑定到Person对象
func (p Person) String() String {
    s := p.Name.Personal + " " + p.Name.Family
    for _, v := range p.Email {
        s += "\n" + v.Kind + ": " + v.Address
    }
    
    return s
}


func main() {
    var person Person
    loadJSON("person.json", &person)
    
    fmt.Println("Person", person.String())
    
}

func loadJSON(fileName string, key interface{}) {

    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := json.NewDecoder(inFile)
    checkError(err)
    inFile.Close()
    
}

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
