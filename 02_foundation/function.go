package main

import "errors"

var zero int;

func main() {
	zero = 3
	println(sum(1,2))
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
