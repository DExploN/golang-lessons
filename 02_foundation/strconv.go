package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	f, _:=strconv.ParseBool("true")
	fmt.Println(f,reflect.TypeOf(f))

	i, _ :=strconv.ParseInt("13", 10 ,64)
	fmt.Println(i,reflect.TypeOf(i))
	

}
