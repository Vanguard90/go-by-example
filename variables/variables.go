package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true // Infer type
    fmt.Println(d)

    var e int // Declare without value
    fmt.Println(e)

    f := "apple" // := is shorthand for declare and assign value
    fmt.Println(f)
}