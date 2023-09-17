package main

import "fmt"

func fibo(n int) {
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = b, a+b
		fmt.Println(i, b)
	}
}

func main() {
	fibo(40)
}
