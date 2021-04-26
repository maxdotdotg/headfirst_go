- goroutines are how go tries to do concurrency
    ```
    go myFunction // now it runs with goroutines?
    ```
- can't use goroutines in functions that use return statements, can't guarantee the goroutine will have completed in time for the return statement
    I think? I don't really get this
    yep, I was correct
- channels allow communication between goroutines
    ```
    // chan is the keyword, and we have to declare the type
    // of values the channel will carry

    var myChannel chan float64 // delaration
    myChannel = make(chan float64) // creation
    myChannel := make(chan float64) // shorthand
    ```
    > Not only do channels allow you to send values from one goroutine to another, they ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it.

- sending data to a channel
    ```
    myChannel <- 3.14
    ```
- receiving data from a channel
    ```
    <- myChannel
    ```
- you know, for a language that reads right-to-left, this feels like a bad choice
- this feels really confusing, it doesn't click for me. I'm gonna try to read this chapter again, or maybe watch some videos about goroutines and channels
- a _very advanced_ video about how channels work under the hood: https://www.youtube.com/watch?v=KBZlN0izeiY
    same presenter for Let's talk about locks! Think it was from Strange Loop?
- channels are FIFO queues used to communicate between goroutines
- > Go makes no guarantees about when it will switch between goroutines, or how long it will keep running one goroutine for. This allows the goroutines to run more efficiently, but it means you can’t count on operations happening in a particular order.

