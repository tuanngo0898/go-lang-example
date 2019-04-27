package main

import ("fmt")

func add(x float64, y float64) float64 {
	return x+y
}

func sub(x, y float64) float64{
	return x-y
}

func mul(x, y float32) float32{
	return x*y
}

func mulRtn(a, b string) (string, string){
	return a, a + " " + b
}

func main(){
	var num1 float64 = 2
	var num2 float64 = 3.1
	var num3,num4 float64 = 2, 3.1
	var num5,num6 float32 = 2, 3.1
	fmt.Println("add: ", add(num1, num2))
	fmt.Println("Sub: ", sub(num3, num4))
	fmt.Println("Mul: ", mul(num5, num6))
	fmt.Println(mulRtn("hello", "worlld"))
	// one line comment
	/*
		Multiple line 
		comment
	*/
	var a int = 62
	var b float64 = float64(a)
	c := a
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}