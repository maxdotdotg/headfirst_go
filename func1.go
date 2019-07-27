package main

import "fmt"

func main() {
    var width, height, area float64
    width = 4.2
    height = 3.0
    area = width * height
    fmt.Printf("%0.2f",area/10.0, "liters of paint\n")
    resultString := fmt.Sprintf("about 1/3: %0.2f\n",1.0/3.0)
    fmt.Println(resultString)
}

//func area(width,height)
