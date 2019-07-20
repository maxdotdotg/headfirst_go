package main

import "fmt"


func main() {
    x := 10
    for x=1; x <=3; x++ {
        fmt.Println(x)
    }
    fmt.Println("outside the loop, x is", x)
}
