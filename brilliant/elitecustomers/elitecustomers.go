package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Customer struct {
	name string
	acc  string
	pref string
}

func main()  {
	f, err := os.Open("./data")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error reading from file")
		return
	}
	custSlice := strings.Split(string(data), "\n")
	var custs []Customer
	for _, c := range custSlice {
		c = strings.TrimLeft(strings.TrimRight(c, "]"), "[")
		s := strings.Split(c, ",")
		a := strings.TrimLeft(strings.TrimRight(s[1], "'"), " '")

		if a >= "A" && a <= "C" {
			cust := Customer{
				name: s[0],
				acc:  a,
				pref: s[2],
			}
			custs = append(custs, cust)
		}
	}
	pref := make(map[string]int)
	for _, c := range custs{
		if _, ok := pref[c.pref]; ok {
			pref[c.pref] += 1
		} else {
			pref[c.pref] = 1
		}
	}
	fmt.Println(pref)
}
