package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	fmt.Println(input)
	
	var name string
    var age int
    fmt.Print("Введите имя: ")
    fmt.Fscan(os.Stdin, &name) 
     
    fmt.Print("Введите возраст: ")
    fmt.Fscan(os.Stdin, &age)
     
    fmt.Println(name, age)
}