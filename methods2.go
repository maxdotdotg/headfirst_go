package main

import (
    "fmt"
    "strings"
)

func main() {
    broken := "G# is pretty #k I guess..."
    replacer := strings.NewReplacer("#","o")

    // call Replace method on the instance of strings.Replacer
    // and pass a string to change
    fixed := replacer.Replace(broken)
    fmt.Println(fixed)
}
