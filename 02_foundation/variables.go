package main

import (
	"fmt"
	"strconv"
	"math"
	"encoding/json"
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

	// array 
	var array [5]string
	array[2] = "privet"
	fmt.Println(array)

	var array2 [2]string = [2]string{"one", "two"}
	fmt.Println(array2)
	fmt.Printf("Types: %T \n", array2) //Types: float64 string boo

	// slice 

	var slice []int
	slice = make([]int,7)
	slice[3] = 6
	slice2 := slice[2:6]
	slice2[2] = 3
	fmt.Println(slice,slice2) // меняются оба слайса

	slice3 := append(slice, 7);
	fmt.Println(slice3)
	slice3[0] = 1 
	fmt.Println(slice, 	slice3)

	var slice4 []int
	slice4 = append(slice4, 2)
	fmt.Println(slice4) 

	// map 

	var firstMap map[string]int
	firstMap = make(map[string]int)
	firstMap["one"] = 1
	println(firstMap["one"])
	if val, ok := firstMap["two"] ; ok {
		println("Exists", val)
	}else{
		println("No exists")

	}

	one, two := firstMap["one"]
	println(one, two)

	secondMap := map[string]int{"name": 123}

	println(secondMap["name"])
	
	delete(secondMap, "name")
	
	println(secondMap["name"])

	// struct

	var firstStruct struct{
		name string
		age int
	}

	firstStruct.age = 123
	firstStruct.name = "Fedor"
	fmt.Println("Struct age: " + strconv.Itoa(firstStruct.age))

	
	var secondStruct secondStructType
	secondStruct.name = "Egor"
	fmt.Println("Struct name: " + secondStruct.name)
	fmt.Println("Struct name from getter: " + secondStruct.getName())
	secondStruct.setName("ne Egor")
	fmt.Println("Struct name after setter: " + secondStruct.getName())


	thirdSecondType := jsonType{Name: "new name", Age: 15}
	res , _:= json.Marshal(thirdSecondType)
	fmt.Println(string(res))
	

	

	var newHellInt hellInt
	newHellInt = 12
	fmt.Println(newHellInt)
}


type secondStructType struct{
	name string
	age int
}

type jsonType struct{
	Name string `json:"name"`
	Age int `json:"fieldNameForAge"`
}

func (r secondStructType) getName () string {
	return r.name
}

func (r *secondStructType) setName (newName string) {
	r.name = newName
}

type hellInt int;