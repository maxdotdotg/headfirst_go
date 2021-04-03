- for maps, keys have to be the same type, and values also have to be the same type
    ```
    var myMap map[string] float64
    // variable myMap is a map that uses strings for keys
    // and float64s for values

    // also, we gotta use make
    var ranks map[string] int
    ranks = make(map[string] int)

    // with literals
    var anotherMap map[string] float64 {"a": 2.0, "b": 3.4, "c": 9.3}
    ```
- maps have the zero-value that corresponds to their types, like slices
- BUT! the zero value for a map itself is nil
    ```
    var nilMap map[int]string
    fmt.Printf("%#v\n", nilMap)
    nilMap[3] = "three"
