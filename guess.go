// player has to guess chosen random number in 10 guesses

package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "time"
    "log"
    "strings"
    "strconv"
)

func main() {
    seconds := time.Now().Unix()
    rand.Seed(seconds)
    target := rand.Intn(100) +1
    fmt.Println("I've picked a number between 1 and 100.")
    fmt.Println("guess the number!")
    fmt.Println(target)

}
