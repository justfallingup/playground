package main

import "fmt"

func row(n int) []int {
	if n == 1 {
		return []int{1}
	}
	prev := row(n-1)
	r := []int{1}
	for i := 0; i < len(prev)-1; i++ {
		r = append(r, prev[i] + prev[i+1])
	}
	r = append(r, 1)
	return r
}

func main()  {
	row23 := row(23)
	sum := 0
	for _, i := range row23 {
		sum += i
	}
	fmt.Println(row23)
	fmt.Println(sum)
}
