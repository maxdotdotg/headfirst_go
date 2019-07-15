// player has to guess chosen random number in 10 guesses

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
    "os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
    //guesses := 10

	fmt.Println("I've picked a number between 1 and 100.")
	fmt.Println("guess the number!")
	fmt.Println(target)

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("what number do you think it is? ")
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

    input = strings.TrimSpace(input)
    guess, err := strconv.Atoi(input)
    if err != nil {
        log.Fatal(err)
    }

    if guess < target {
        fmt.Println("Oop! guess higher!")
        //fmt.Println("you have", guesses, "left")
    } else if guess > target {
        fmt.Println("Oop! guess lower!")
        //fmt.Println("you have", guesses, "left")
    } else {
        fmt.Println("you got it!")
    }
}
