package main

func main() {
	for x := 0; x < 10; x += 2 {
		println(x)
	}
	x := 0
	for x < 10 {
		println(x)
		x += 2
	}

	var array2 [2]string = [2]string{"one", "two"}

	for i := 0; i < len(array2); i++ {
		println(array2[i])
	}

	for index, value := range array2 {
		println(index, value)
	}

	for index := range array2 {
		println(index)
	}
}
