- arrays are lists but all the same type and have a static length
- `var myArray [4]string`: myArray has 4 elements, all of type string
- assignment:
    ```
    var notes [7]string
    notes[0] = "do"
    notes[1] = "ray"
    notes[2] = "me"
    ```
- usage: `fmt.Println(notes[0])`
- zero/empty values in the array have the zero value of the type
- individual elements of the array can be updated
    ```
    var counters [3]int
    counters[0]++ // increment first element from 0 to 1
    counters[0]++ // increment first element from 1 to 2
    counters[2]++ // increment THIRD element from 0 to 1
    fmt.Println(counters[0], counters[1], counters[2]) // returns 2 0 1
    ```
- "array literal", defining an array where you know types and values: `var newArray 3[int] = [3]int {1, 2, 3}`
    feels a bit awkard, re-stating the type definition... not sure why that's the way?
- shorthand variable declaration works for array literals too: `primes := 5[int] {1, 2, 3, 5, 7}`

- REMINDER: debugging with `%#v` works with `fmt.Printf`, NOT `fmt.Println`

- I like the `range` keyword, so much easier to use than the c-style loop
    ```
    for index, value := range array {
        // do things with both vars her
    }
    ```
- regarding `go install`: as of go 1.16, modules are on by default, which breaks the `go install` comamnd from the book:
    > go install github.com/headfirstgo/average
    the workaround for this is to set `GO111MODULE=off` when running `go install`
    https://blog.golang.org/go116-module-changes
- similar for `go run` after I upgraded, which may not have been so smart... now I need `GO111MODULE=off` when I use custom libs (libs from `~/go/src`) and that's... not great? I feel like I messed something up here, IDK



