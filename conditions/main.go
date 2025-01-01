package main

import "fmt"

func main() {
	fmt.Println("Hello world!!")

	arr := [5]string{"Go", "Python", "C++"}

	for _, item := range arr {

		fmt.Println(item)

	}
}
