package main

import "fmt"

func main() {
    // string and int formatting
    fmt.Printf("the %s cost %d cents each.\n","gumballs",23)

    // float formatting with 2 decimal places
    fmt.Printf("that'll be %0.2f please.\n",0.23*5)

    // float formatting on its own
    fmt.Printf("that'll be %f please.\n",0.23*5)
}
