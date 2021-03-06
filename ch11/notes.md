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

- > When you have a variable of an interface type, the only methods you can call on it are those defined in the interface

- type assertions are when you define a variable using an interface and need to access methods associated with the concrete type.
    ```
    // noiseMaker is type NoiseMaker (an interface) that
`   // type Robot satisfies
    var noiseMaker NoiseMaker = Robot("company")

    // robot is type Robot, which satisfies the NoiseMaker
    // interface, and this assertion gives us access to
    // the other Robot methods
    var robot Robot = noiseMaker.(Robot) // DOT PARENTHESES CONCRETE TYPE
    ```
    I'm really fuzzy on this

- if a variable is defined as an interface type, then it's limited to the methods associated with the interface. to give the variable access to methods associated with the concrete type, it needs to be converted using a type assertion
- `ok` will catch whether or not the type assertion (and I keep wanting to say conversion) is successful, so conditionals based on `ok` can be used
    ```
    robot, ok := noiseMaker.(Robot)
    if ok != nil {
        // handle error
    }
    ```

- the error interface
    ```
    type error interface {
        Error string
    }
    ```
    we can create custom errors that use the same interface, see [errors.go](errors.go)
- also, the `error` interface doesn't need to be imported, it's "part of the 'universe block'", or what I'm used to thinking of as built-ins
- the stringer interface
    ```
    type Stringer interface {
        String() string
    }
    ``` 
    it allows "any type to decide how it will be displayed when printed", see [stringers.go](stringers.go)

- the empty interface (used by fmt.Println(), for example)
    ```
    type Something interface {}
    ```
    > The empty interface doesn’t have any methods that are required to satisfy it, and so every type satisfies it.
    > If you declare a function that accepts a parameter with the empty interface as its type, then you can pass it values of any type as an argument
    > [The] empty interface doesn’t have any methods. That means there are no methods you can call on a value with the empty interface type!
