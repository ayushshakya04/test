package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointers")

	// 	// taking a normal variable
	// 	var x int = 23

	// 	//declare of pointer
	// 	var p *int

	// 	//initialization of pointer
	// 	p = &x

	// 	//display the result
	// 	fmt.Println("value stored in x= ", x)
	// 	fmt.Println("value stored in x= ", &x)
	// 	fmt.Println("value stored in p= ", *p) // normal *p gives you value that is assign to the p.
	// 	fmt.Println("value stored in p=", p)   // p gives you the memory address.
	// }
	myNumber := 23
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	// *ptr = *ptr * 2
	// fmt.Println(myNumber)
	// fmt.Println(myNumber)
}
