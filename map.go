package main

import "fmt"

func main(){
	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jess"]  = 50
	grades["Sam"]	= 80
	fmt.Println(grades)
	fmt.Println(grades["Jess"])
	delete(grades, "Jess")
	fmt.Println(grades)
	fmt.Println(grades["Jess"])

	for k, v := range grades{
		fmt.Println(k,":", v)
	}
}