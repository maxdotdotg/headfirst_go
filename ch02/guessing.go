// guess a number between 1 and 100
// 10 guesses, report remaining before each guess

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
	// seed the RNG
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	//fmt.Println(target)

	// start success as false, change only on victory condition
	success := false

	// loop over guesses, hardcoded to 10
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("guesses left:", 10-guesses)

		// read and parse input
		fmt.Print("Guess the number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		// compare guess to target
		if guess > target {
			fmt.Println("too high!")
		} else if guess < target {
			fmt.Println("too low!")
		} else {
			fmt.Println("you got it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("sorry you didn't get it, it was", target)
	}

}
