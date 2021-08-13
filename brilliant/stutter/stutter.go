package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	b := make([]byte, 16)
	var d []byte
	f, err := os.Open("./data")
	if err != nil {
		fmt.Println("error open file")
		return
	}
	defer f.Close()
	for {
		n, err := f.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading from file")
		}
		b = b[:n]
		d = append(d, b...)
	}
	sl := strings.Split(string(d), "\n")

	var clear []string
	clear = append(clear, sl[0])
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] != sl[i+1] {
			clear = append(clear, sl[i+1])
		}
	}
	fmt.Println(clear)
	fmt.Println("чистый список", len(clear))
}
