package main

import (
	"fmt"
	"reflect"
)

func main() {
	var quantity int
	fmt.Println(&quantity)
	fmt.Println(reflect.TypeOf(&quantity))

	var quantityRef *int
	quantityRef = &quantity
	fmt.Println(*quantityRef)
	fmt.Println(quantity)

	quantity = 2 

	fmt.Println(*quantityRef)
	fmt.Println(quantity)

	*quantityRef = 3

	fmt.Println(*quantityRef)
	fmt.Println(quantity)

	fmt.Println(quantityRef)
	fmt.Println(&quantity)

	changeRef(&quantity)

	fmt.Println(*quantityRef)
	fmt.Println(quantity)
}

func changeRef(ref *int){
	*ref = 1000
}
