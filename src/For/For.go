//source: https://gobyexample.com/for
//Shows examples of the different use of for loops

package main

import "fmt"

func main() {
	i := 1
    //basic loop
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    //standard loop
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    //this will loop forever or until it is broken using the break command or return
    for {
        fmt.Println("loop")
        break
    }
}