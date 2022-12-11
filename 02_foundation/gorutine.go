package main

import (
	"fmt"
)

type Info struct {
	text string
}

func main() {
	 result:= make(chan int)
	 go summ(1,2, result)
	 fmt.Println(<-result)

	 infoResult := make(chan Info)
	 go getInfo(infoResult)

	 info:= <-infoResult
	 fmt.Println(info.text)
}

func summ(a int , b int, result chan int )  {
	result <- a+b
}

func getInfo(result chan Info){
	info := Info{text: "Lorem ipsum"}
	result<- info
}
