package main

import "fmt"

const tempebulK float64 = 373  //ponto de obulição da agua em kelvin

func main() {

	var c float64 = tempebulK - 273
	fmt.Println("A temperatura de ebulição da agua em ºC é de:", c)

}
