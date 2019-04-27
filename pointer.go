package main

import "fmt"

func main(){
	x:= 15
	a:= &x
	fmt.Println(x)
	fmt.Println(*a)
	fmt.Println(a)
	fmt.Println(&x)

	*a = 16
	fmt.Println(*a)
	fmt.Println(x)
}