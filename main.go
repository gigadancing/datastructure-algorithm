package main

import "fmt"

func foo(arr [4]int) {
	arr[0] = 100
}

func main() {
	a := [4]int{1, 2, 3, 4}
	foo(a)
	fmt.Println(a)
}
