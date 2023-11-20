package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n // arithmetic with arbitrary precision
    fmt.Println(d)

    fmt.Println(int64(d)) // no type until explicit conversion

    fmt.Println(math.Sin(n)) // can be given type due to context
}