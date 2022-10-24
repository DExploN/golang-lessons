package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		println(err)
	}
}
