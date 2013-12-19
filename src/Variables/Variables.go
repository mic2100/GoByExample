//Source: https://gobyexample.com/variables
//A basic example of how to assign data types to variables

package main

import "fmt"

func main() {
	var a string = "initial" //this value must be a string
    fmt.Println(a)

    //allows multiple vars to be set at once
    var b, c int = 1, 2 //values b & c must both be integers
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    //sets the var to the zero-value, in this case it is an int (0)
    var e int
    fmt.Println(e)

    //the := auto assigns the data type to the f var to string,
    //if anything other than a string is set an error will be produced
    f := "short"
    fmt.Println(f)
}
