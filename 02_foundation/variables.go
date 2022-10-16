package main

import (
	"fmt"
	"math"
)

func main() {
	// Объявление переменной
	var quantity int;
	quantity = 1;
	fmt.Println(quantity)

	// Объявление с одновременным присвоением
	quantity2 := 2
	fmt.Println(quantity2)

	// Присвоение двух переменных сразу
	quantity, quantity2 = 3,4
	var quantity3 = 5;
	fmt.Println(quantity, quantity2, quantity3)
	
	var length ,height int = 1,2
	fmt.Println(length, height)


	// Преобразования типов

	intVar:= 12
	var floatVar float32 = 1.5
	
	fmt.Println("Преобразования типов")
	fmt.Println(intVar * int(floatVar))
	fmt.Println(math.Floor(float64(floatVar)))
	fmt.Println(float64(intVar) * math.Floor(float64(floatVar)))
	fmt.Println(float32(intVar) * floatVar)

}