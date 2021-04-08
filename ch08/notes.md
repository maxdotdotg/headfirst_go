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
        field1 string
        field2 int
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

    it's cheaper to pass pointers than to pass structs, so do it that way... I guess?

- structs can be exported as well, and must be Capitalized to do so, and so must their attributes...
- and of course, struct literals
    ```
    # referencing myStruct above
    testingStruct := myStruct{field1: "value for field1", field2: 9001}
    ```
- structs can be nested
    ```
    type myType struct {
        field1 string
        field2 int
        // field3 MyOtherStruct // accessed "normally"?
        MyOtherStruct // access fields anonymously
    }
    
    s := myStruct.MyOtherStruct
    s.MyOtherStruct.OtherStructField1 = "value of some attribute in MyOtherStruct"
    ```
    
    and nested structs can be accessed more easily through anonymous fields, see `[structs_testing/main.go](structs_testing/main.go)` for examples of usage, and `[magazine.magazine.go](magazine/magazine.go)` for nested struct definitions
 
