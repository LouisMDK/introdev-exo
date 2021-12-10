package main

import "fmt"
import "strconv"

func main() {
	fmt.Println(q20())
}

func q20() int {
	var s int = factorielle(100)
	fmt.Println(s)
	return s
}

func somme_digits(n int) (res int) {
	var ns string = strconv.Itoa(n)
	for i := 0; i < len(ns); i++ {
		b, _ := strconv.Atoi(string(ns[i]))
		res += b
	}
	return res
}

func factorielle(n int) (res int) {
	res = 1
	for i := 2; i <= n; i++ {
		fmt.Println(res)
		res *= i
	}
	return res
}