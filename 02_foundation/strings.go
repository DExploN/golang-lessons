package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Replace substring");
	broken := "from var";
	replacer := strings.NewReplacer("from", "to");
	fmt.Println(replacer.Replace(broken));

	// Количество символов
	var s string
	s = "Privet"
	fmt.Println(len([]rune(s)))

	// Удаление пробелов и переводов строк
	s = "\t trimed \n \n \n";
	s = strings.TrimSpace(s);
	fmt.Println(s)
}
