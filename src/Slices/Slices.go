//source: https://gobyexample.com/slices
//examples of the slices array data type

package main

import "fmt"

func main() {
	//makes a new slices of strings with 3 elements
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	//they can be set just like normal arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	//values can easily be appended
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	//makes a new slice with the same length as `s`
	c := make([]string, len(s))
	
	//copies the content from `s` to `c`
	copy(c, s)
	fmt.Println("cpy: ", c)
	
	//this will extract elements 2 -> 4
	l := s[2:5]
	fmt.Println("sl1:", l)

	//this will extract all element <5
	l = s[:5]
	fmt.Println("sl2:", l)
	
	//this will extract all elements >=2
	l = s[2:]
	fmt.Println("sl3:", l)

	//declare and init a slice in one row
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//a multi-dimensional slice 
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
