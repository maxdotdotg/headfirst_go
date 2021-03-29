- slices are mutable arrays! thank gods
- `var mySlice []string`, empty brackets because it's mutable
- gotta _make_ a slice
    ```
    var notes []string
    notes = make([]string 7) // make a slice with 7 empty strings)
    // or shorthand
    notes := make([]string 7)

    // and slice literals
    otherSlice := []string {"hi", "hello", "so on"}

    ```
- slices are built on top of arrays.
    not sure I care a whole lot, though it's nice to know about a fixed-size list vs a flexible list, and you can make a slice from an array
    ```
    underlyingArray := [5]string {"a", "b", "c", "D", "e"}
    slice := underlyingArray[:3]
    ```
- and of course, you can change things out from under yourself!
> because a slice is just a view into the contents of an array, if you change the underlying array, those changes will also be visible within the slice!
    ch06
- append! like I'd expect!
    ```
    slice := []string {"A", "b"}
    fmt.Println(slice) // A b
    slice = append(slice, "c")
    fmt.Println(slice) // A b c
    ```
- if you're gonna append new values to your slice, do so with the same variable name, otherwise you end up with inconsistent behavior, since underlying arrays can't be modified, you don't know what will be changed and what won't (see `slices1.go` for an example)
- when slices are made with `make`, they take the zero value of their type 
    ```
    // zero value of an empty string
    notes = make([]string 7) 

    // zero value of nil
    // `fmt.Printf("%#v", notes2)` shows `[]string(nil)`
    // called a "nil slice"
    var notes2 = []string 
    ```
- fun things about paths: go apparently doesn't expand bash env vars like `$HOME` or the home shortcut `~`, so I have to pass the absolute path to `datafile.GetFloats()` if I want it to work. This feels kinda silly, like there's something I'm missing, but it's what I got...
    ```
    ➜  ch05 git:(main) ✗ go run file.go
    2021/03/28 20:28:10 open ~/go/head-first-go/ch05/data.txt: no such file or directory
    <nil>
    &{0xc0000ba120}
    ```
- variadic functions: functions that take any number of params, like `fmt.Println()`
    ```
    func newFunc(param1 int, param2 ... string) {
    // do things here
    }
    ```
    yes, the elipses are actually required
    > use an ellipsis (...) before the type of the last (or only) function parameter in the function declaration. ... The last parameter of a variadic function receives the variadic arguments as a slice, which the function can then process like any other slice.
    ch06

    ```
    func severalInts(numbers ...int) {
        fmt.Println(numbers)
    }
    func main() {
        severalInts(1) // prints [1]
        severalInts(1, 2, 3) // prints [1 2 3]
    }
    ```

- I don't really get how this works. I know it _does_, but I don't get why. Is the content/value of the slice being passed?
    > add an ellipsis (...) following the slice you want to use in place of variadic arguments.
    > When calling a variadic function, you can use a slice in place of the variadic arguments by typing an ellipsis after the slice: `inRange(1, 10, mySlice...)`
