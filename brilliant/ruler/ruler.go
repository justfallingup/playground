package main

import (
	"fmt"
	"math"
)

func rule(l, r, h int)  {
	if h > 0 {
		x := (l + r)/2
		fmt.Println("mark ", x, " height ", h)
		rule(l, x, h-1)
		rule(x, r, h-1)
	}
	return
}

func main()  {
	rule(0, int(math.Pow(2,4)), 4)
}
