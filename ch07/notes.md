- for maps, keys have to be the same type, and values also have to be the same type
    ```
    var myMap map[string] float64
    // variable myMap is a map that uses strings for keys
    // and float64s for values

    // MUST use _make_, otherwise it's a nil map
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
    nilMap[3] = "three" // panic, throws "assignment entry in nil map"
    ```
- zero values vs assigned values
    > accessing a map key optionally returns a second, Boolean value. It will be true if the returned value has actually been assigned to the map, or false if the returned value just represents the default zero value.
    ch07
- delete items from maps, `delete(myMap, "some key I told to GTFO")`, no assignment like `append`, curious. Key doesn't have to be a string, just used it for an example, see [`delete_from_map.go`](delete_from_map.go) for more.
- looping through maps is done the same way as with arrays
    ```
    for k, v := range myMap {
        // do some stuff
    }
    ```
- and it turns out, if you only want _keys_, you can skip the values: `for k := range myMap` etc etc. Note that this does NOT work if you only want the _values_, you still have to use the underscore to skip `for _, v := range myMap`
