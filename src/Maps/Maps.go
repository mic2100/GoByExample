//source: https://gobyexample.com/maps
//an example of Go's associative Map datatype

package main

import "fmt"

func main() {
	//make a map with string keys containing ints
	m := make(map[string]int)
	
	//assign key/values to the map, the same sort of thing as you do with an array 
	m["k1"] = 7
	m["k2"] = 13
	
	fmt.Println("map: ", m)
	
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))
	
	//delete an element from the map
	delete(m, "k2")
	fmt.Println("map: ", m)
	
	_, prs := m["k2"]
	fmt.Println("psr: ", prs)
	
	//create a map and assign values in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
