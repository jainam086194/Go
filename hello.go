/*
	 program to print
		hello world to the screen

author: Jainam Shah
date: 26/12/2024
*/
package main

import "fmt"

func main() {
	//fmt.Println("HEllo World")
	//var my_int int = 200
	//var my_float float32 = 343.233
	//var my_string string = "Jainam Shah"
	//var my_bool bool = true
	//fmt.Println(my_int)
	//fmt.Println(my_float)
	//fmt.Println(my_string)
	//fmt.Println(my_bool)
	//
	//x := "Jainam"
	//fmt.Println(x)
	//x = "100"
	//fmt.Println(x)
	//var a int
	//var b string
	//var c bool
	//a = 20
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//var a, b, c, d = 1, 3, 5, 7
	//fmt.Println(a, b, c, d)
	//var (
	//	name         string = "Jainam"
	//	id           int    = 29
	//	pass_or_fail bool   = False
	//)
	//
	//fmt.Println(name, id, pass_or_fail)
	//const CONSTANT_VAR float32 = 203.232324323423
	////CONSTANT_VAR = 343.3
	//fmt.Println(CONSTANT_VAR)
	const (
		CONST1 int     = -300
		CONST2 float32 = 34.333342424234242342343
		CONST3 bool    = false
		CONST4 string  = "Hi World"
	)

	fmt.Println(CONST1, "\n", CONST2, "\n", CONST3, "\n", CONST4)
	fmt.Printf("Value of CONST1 is: %+d and Type of CONST1 is %T\n", CONST1, CONST1)
	fmt.Printf("Value of CONST2 is: %g and Type of CONST2 is %T\n", CONST2, CONST2)
	fmt.Printf("Value of CONST3 is: %t and Type of CONST3 is %T\n", CONST3, CONST3)
	fmt.Printf("Value of CONST4 is: %q and Type of CONST4 is %T\n", CONST4, CONST4)

	var x int8 = 100
	fmt.Printf("Type: %T, value: %v\n", x, x)

	var (
		arr1 = [...]int{1, 2, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 100, 200, 1000000}
		cars = [3]string{"BMW", "Mercedes", "Nissan"}
		arr2 = [10]float32{4: 23.2323, 7: 100.23232}
	)
	fmt.Println(len(arr1))
	fmt.Println(cars)
	fmt.Println(len(arr2))

}
