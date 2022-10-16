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
}
