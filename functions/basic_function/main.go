package main

import "fmt"

func getNumber() int {
<<<<<<< HEAD
	return 100
=======
	return 20000
>>>>>>> functions
}

func main() {
	num := getNumber()
	fmt.Println(num)
}
