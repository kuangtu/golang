package main

import (
    _ "fmt"
)

type MyInt int

type IntAlias = int

func main() {
    var a int = 1
    var b MyInt = 2
    c := a + int(b)
    var d IntAlias = 1
    var e int = 2
    f := d + e
}
