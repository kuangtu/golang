package main

import (
    "fmt"
    "encoding/json"
    "os"
    "io/ioutil"
)

type MyStructList struct {
    StructList []MyStruct `json:"structlist"`
}

type MyStruct struct {
    Name string `json:"Name"`
    Age int `json:"Age"`
}

func main() {

    stud := MyStruct {
        Name: "tony",
        Age: 13}
    
    data, err := json.Marshal(stud)
    if err != nil {
        fmt.Println("json marshal failed:%s", err)
    }
    
    fmt.Println(string(data))
    
    
    //读取json file
    filepath := "conf.json"
    jsonfile, err := os.Open(filepath)
    if err != nil {
        fmt.Println("open error.")
        os.Exit(1)
    }
    defer jsonfile.Close()
    
    //读取json文件
    byteBuffer, err := ioutil.ReadAll(jsonfile)
    fmt.Println(string(byteBuffer))
    
    var mylists MyStructList
    
    unerr := json.Unmarshal(byteBuffer, &mylists)
    
    if unerr != nil {
        fmt.Println(unerr)
        os.Exit(1)
    }
    //输出json信息
    listlen := len(mylists.StructList)
    fmt.Println("the json list len is:", listlen)
    for i := 0; i < listlen; i++ {
        fmt.Println("Name: ", mylists.StructList[i].Name)
        fmt.Println("Age: ", mylists.StructList[i].Age)
        fmt.Println("---")
    }
    
    
}