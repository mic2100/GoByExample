//source: https://gobyexample.com/arrays
//Simple example showing how to manage simple array structures

package main

import "fmt"

func main() {
	//create an array with 5 empty keys
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }

	fmt.Println("2d: ", twoD)
}
