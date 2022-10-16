package main

import (
	"fmt"
	"math"
	"strings"
	"reflect"
)

func main() {
	name := "Go Developer"
	fmt.Println("Azure for", name)
	fmt.Println(math.Floor(2.75))
	// rune. Руны через одиночные кавычки, показывают номер символа
	fmt.Println('a')
	fmt.Println(strings.Title("string to upper"))
	// Рефлексия типов
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(15))


}
