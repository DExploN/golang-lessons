package main

import (
	"log"
)

func main() {
	text := "log message"
	loger:=log.Default()
	loger.Println("print");
	log.Fatalln("fatal", text)
	// no execute. Exit with status 1 after FatalLn
	log.Panicln("panic", text)
}