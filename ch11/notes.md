- > [Interfaces] let you define variables and function parameters that will hold any type, _as long as that type defines certain methods_. [emphasis mine]

    kinda like a contract based on functions 

- interface definition
    ```
    type myInterface interface {
        methodWithoutParams()
        methodWithParams(float64)
        methodWithRreturnValue() string
    }
    ```

    if a type has all these methods, then it *satisfies* the interface
    and a variable can be created with the interface type, so it can be any type that satisfies the interface (see [interfaces1.go](interfaces1.go) for an example)
    BUT! A variable that's an interface CANNOT use methods outside of the interface

- > A concrete type specifies not only what its values can do (what methods you can call on them), but also what they are: they specify the underlying type that holds the value’s data.
    > Interface types don’t describe what a value is: they don’t say what its underlying type is, or how its data is stored. They only describe what a value can do: what methods it has.

- 
