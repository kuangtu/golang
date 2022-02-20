
package main

import "fmt"


const ( 
    mutexLocked = 1 << iota
    mutexWoken
    mutexStarving
    mutexWaiterShift = iota
    starvationThresholdNs = 1e6
)


func main() {
    fmt.Println(mutexLocked)
}
