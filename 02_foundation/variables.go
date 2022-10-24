package main

import (
	"fmt"
	"math"
)

func main() {
	// Объявление переменной
	var quantity int
	quantity = 1
	fmt.Println(quantity)

	// Объявление с одновременным присвоением
	quantity2 := 2
	fmt.Println(quantity2)

	// Присвоение двух переменных сразу
	quantity, quantity2 = 3, 4
	var quantity3 = 5
	fmt.Println(quantity, quantity2, quantity3)

	var length, height int = 1, 2
	fmt.Println(length, height)

	// Преобразования типов

	intVar := 12
	var floatVar float32 = 1.5

	fmt.Println("Преобразования типов")
	fmt.Println(intVar * int(floatVar))
	fmt.Println(math.Floor(float64(floatVar)))
	fmt.Println(float64(intVar) * math.Floor(float64(floatVar)))
	fmt.Println(float32(intVar) * floatVar)

	// Если хоть одна переменная в множественном присвоении новая, то можно

	a, b := 1, 2
	b, c := 3, 4
	fmt.Println(a, b, c)


	var array [5]string
	array[2] = "privet"
	fmt.Println(array)

	var array2 [2]string = [2]string{"one", "two"}
	fmt.Println(array2)
	fmt.Printf("Types: %T \n", array2) //Types: float64 string boo

}
