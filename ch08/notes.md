- structs! many data types in one structure! yay.
    ```
    var myStruct struct {
        word string
        quantity int
    }
    ```
- access fields with `.`: `fmt.Println(myStruct.word`)
- assign fields with `.`: `myStruct.quantity = 5`
- type definition (kinda like classes, I guess? IDK)
    ```
    type myType struct {
        // fields go here
    }

    var instanceOfMyType myType
    ```
- functions can take custom types as well
- to update existing assigned values in a struct when using a function, use pointers
    ```
    newFunc(&myStruct)
    ...
    ...
    ...
    func newFunc(s *myStruct){
        s.field = "some new value"
    }
    ```
- to access a pointer to a field in a struct: `(*myStruct).myField` ? I think? pointers still confuse the crap out of me
- to access a field in a struct that _is_ a pointer: `*myStruct.myField`
- > Functions receive a copy of the arguments they’re called with, even if they’re a big value like a struct. That’s why, unless your struct has only a couple small fields, it’s often a good idea to pass functions a pointer to a struct, rather than the struct itself. 



