package main

import "fmt"

func present(l []int, n int) bool {
	for _, i := range l {
		if n == i {
			return true
		}
	}
	return false
}

func unique(l []int) []int {
	var un []int
	for _, i := range l {
		if present(un, i) == false {
			un = append(un, i)
		}
	}
	return un
}

func main()  {
	l := []int{2, 7, 11, 7, 2}
	fmt.Println(present(l, 2))
	fmt.Println(present(l, 42))
	fmt.Println(unique(l))

}
