package main

import "fmt"

func mcCathy91(n int) int {
	if n > 100 {
		fmt.Println(n)
		return n - 10
	}
	return mcCathy91(mcCathy91(n + 11))
}

func main()  {
	m := mcCathy91(98)
	fmt.Println("result", m)
}
