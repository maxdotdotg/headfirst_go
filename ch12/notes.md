- `defer` will run a given line of code when the function exits
- only function calls and method calls can be `defer` d
- `panic`, we can call it and it'll, you guessed it, panic, dump a stack trace, and exit
- when `panic` gets called, all `defer` d functions run before exiting
- `recover` handles `panic`, and can allow the program to still exit with status 0, even when `panic` gets called
- any code below a `panic` will never get called, even when using `recover`; the function that contains `panic` exits at that point, `recover` picks up after that function exits.

