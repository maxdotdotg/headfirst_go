package main

import "fmt"

func main() {
    fmt.Println("loop one, I think it'll be 3, 2, 1")
    for x :=3; x >=1; x-- {
        fmt.Println(x)
    }

    fmt.Println("loop two, I think it'll be 1,2")
    for x := 1; x < 3; x++ {
        fmt.Println(x)
    }

    fmt.Println("loop three, I think it'll be 3,4++forever")
    for x :=1; x <= 3; x++ {
        fmt.Println(x)
    }
    fmt.Println("I was wrong! it never triggers, because it's always false")
    fmt.Println("I was wrong because I made a typo! > instead of <\n\n")

    fmt.Println("loop four, I think it'll be nothing")
    for x :=2; x <=3; x++ {
        fmt.Println(x)
    }
    fmt.Println("I was wrong again. I misread the condition.\n")

    fmt.Println("loop five, it should print 1, 3")
    for x := 1; x <= 3; x+=2 {
        fmt.Println(x)
    }


}
