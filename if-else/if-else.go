package main

import "fmt"

func main() {

	// No paranthesis needed around conditionals

    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if 7%2 == 0 || 8%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

    if num := 9; num < 0 { // You can declare variables in a statement
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// There is no ternary if/else (a ? b : c) in Go!