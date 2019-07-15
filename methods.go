package main

import (
    "fmt"
    "time"
)

func main() {
//    now := time.Now()
//    year := now.Year()

    //now is a variable that holds a time.Time value
    var now time.Time = time.Now() //calling methods on a type

    // call the Year method on the time.Time value
    // methods are type-dependant, which makes sense, I guess?
    var year int = now.Year()
    var hour int = now.Hour()

    fmt.Println(year)
    fmt.Println(hour)
}
