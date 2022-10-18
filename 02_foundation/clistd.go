package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(input)
	}

	var name string
    var age int
    fmt.Print("Введите имя: ")
    fmt.Fscan(os.Stdin, &name) 
     
    fmt.Print("Введите возраст: ")
    fmt.Fscan(os.Stdin, &age)
     
    fmt.Println(name, age)
}