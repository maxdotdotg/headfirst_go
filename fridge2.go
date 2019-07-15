package main

import (
//    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
//    fmt.Print("which file? ")
//    reader := bufio.NewReader(os.Stdin)
//    targetFile, err := reader.ReadString('\n')
//    if err != nil {
//        log.Fatal(err)
//    }

//    fileInfo, err := os.Stat(targetFile)
    fileInfo, err := os.Stat("my.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(fileInfo.Size())
}
