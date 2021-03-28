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
