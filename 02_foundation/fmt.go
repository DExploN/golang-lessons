package main

import "fmt"

func main() {
	// 0 это минимльное количество символов перед точкой. Если число больше чем результат - то дополняется пробелами.
	fmt.Printf("About one-third: %10.2f\n", 1.0/3.0)
	// Вторая - число знаков после точки, заполняется 0 если не хватает в ответе
	resultString := fmt.Sprintf("About one-third: %0.5f\n", 100.0/5.0)
	fmt.Printf(resultString)


	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	// подходящее значение само выбирается
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	// как выглядит в golang  
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true) // Values: 1.2 "\t" true
	// выводит типы
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true) //Types: float64 string boo
	fmt.Printf("Percent sign: %%\n")
}
