package main

import (
	"fmt"
)

func main() {
	var text string
	// Необходима инициализация, так как если инициализировать внутри if,
	// то text будет недоступен в printLn из-за области видимости
	if true {
		text = "success"
	} else {
		text = "fail"
	}
	fmt.Println(text)
}
