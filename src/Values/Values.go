//Source: https://gobyexample.com/values
//A basic example showing how different data types are handled

package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") //concatenating strings

	fmt.Println("1 + 1 = ", 1+1)         //integer addition
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0) //float divisions

	fmt.Println(true && false) //prints boolean false
	fmt.Println(true || false) //prints boolean true
	fmt.Println(!true)         //prints boolean false
}
