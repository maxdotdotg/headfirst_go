package main

import "fmt"

func main() {
    // declare myInt
    // declare myIntPointer as pointer to int
    // assign val to myInt
    // assign pointer to myIntPointer
    // print val of myIntPointer

    var myInt int
    var myIntPointer = &myInt
    myInt = 42
    fmt.Println(*myIntPointer)
}
