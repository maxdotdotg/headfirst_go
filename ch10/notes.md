- > Setter methods are methods used to set fields or other values within a defined type’s underlying value.
By convention, Go setter methods are usually named in the form SetX, where X is the thing that you’re setting.

- setter and getter methods allow access to values in struct fields which aren't exported, this means we can have some validation up-front with setters. For example, setter and getter methods in [calendar/date.go](calendar/date.go) are used in [dates2.go](dates2.go)

- > But if any method on a type takes a pointer receiver, convention says that they all should, for consistency’s sake.

    not sure how I feel about that, but OK, I guess?

- encapsulation (controlling access to underlying fields and structs and stuff) gets used when there's a need for validation on those fields, like we did for months in a year

