- custom types, or defined types, are usually built on structs, but can be built on any type: `type Gallons float64`
- > If you have a variable that uses a defined type, you cannot assign a value of a different defined type to it, even if the other type has the same underlying type. This helps protect developers from confusing the two types. ... But you can convert between types that have the same underlying type.

- > defined types cannot be used in operations together with values of a different type, even if the other type has the same underlying type. 

- receiver methods! functions that are availale _only_ to a given type (and different types can have receivers with the same name), accessed via dot-methods
    ```
    type MyType string


    // m is the receiver parameter name
    // MyType is the receiver parameter type
    func (m MyType) sayHi() {
        fmt.Println("Hi from", m)
    }

    value := MyType("a MyType value")
    value.sayHi()
    ```
