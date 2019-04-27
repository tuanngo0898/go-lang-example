package main

import "fmt"

const usixteenbitmax = 65535
const kmh_multiple float64 = 1.60934

type car struct{
	gas_pedal uint16
	brake_pedal uint16
	stear_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)/kmh_multiple
}

func (c *car) newTopSpeed(new_top_speed float64){
	c.top_speed_kmh = new_top_speed
}

func newerTopSpeed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main(){
	a_car := car{gas_pedal: 200, brake_pedal: 0, stear_wheel: 12, top_speed_kmh: 200.5}

	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())


	a_car.top_speed_kmh = 100
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.newTopSpeed(300)
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car = newerTopSpeed(a_car,400)
	fmt.Println(a_car.top_speed_kmh)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}