package main

import "errors"

var zero int;

func main() {
	zero = 3
	println(sum(1,2))

	println(sumManyIntegers(1,3,5))

	arr:= []int{1,2,3,4,5}
	println(sumManyIntegers(arr...))
}

func sum(first int, second int) int {
	return first + second + zero
}

func cube(first int)(int, error){
	result := first * first 
	if result == 0{
		return 0, errors.New("zero")
	}else{
		return result, nil
	}
}


func sumManyIntegers(integers ...int) int {
	result := 0
	for _, integers:= range integers{
		result+=integers
	}
	return result
}