package main

import "fmt"

type car struct{
	gas_pedal uint16
	brake_pedal uint16
	stear_wheel int16
	top_speed_kmh float64
}

func main(){
	a_car := car{gas_pedal: 1234, brake_pedal: 0, stear_wheel: 12, top_speed_kmh: 200.5}
	fmt.Println(a_car)
}