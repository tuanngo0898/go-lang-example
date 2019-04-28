package main
import "fmt"

func main(){
	for i:= 1; i<10; i++{
		fmt.Println(i)
	}
	j := 0
	for j<10{
		fmt.Println(j)
		j++
	}
	x:= 5
	for{
		fmt.Println("Do stuff",x)
		x+= 3;
		if x > 25{
			break
		}
	}
}