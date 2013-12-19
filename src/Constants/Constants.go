//source: https://gobyexample.com/constants
//Shows how to use constants

package main

import (
    "fmt"
    "math"
)

//constants defined outside of the methods are globally available
const s string = "constant"

func main() {
	fmt.Println(s)

    //constants defined inside a method are only available iits scope
    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)

    //Casts int64 against the `d` constant
    fmt.Println(int64(d))

    //Caclulates the sin value, expects a float64 to be supplied
    fmt.Println(math.Sin(n))
}
