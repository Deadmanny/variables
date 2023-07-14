package main

import "fmt"

func variables() {

	// these are different ways how variables can be declared n intialized
	// uncomment only one at once

	var a string = "hello"
	//var a = "bey"
	//a := "fine" // short hand declaration n intialization
	fmt.Println(a)

	var b int = 3
	//var b = 4
	//b := 5
	//var b, c int = 4, 5
	//var b, c = "k", "k"
	//b = "j"
	//c = "p"
	//fmt.Println(b, c)
	fmt.Println(b)

	var arr = [5]int{1, 2, 3, 4, 5}
	//arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr)

	var d bool = true
	//d := !false
	fmt.Println(d)

}
func main() {
	variables()

}
