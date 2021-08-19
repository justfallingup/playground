package main

import (
	"fmt"
	"playgroung/freeCodeCamp/datastrustures/stack"
	"strings"
)

func valid(s string) bool {
	stackBraces := stack.New()
	braces := strings.Split(s, "")

	bracesPairs := map[string]string{
		"(":")",
		"[":"]",
		"{":"}",
	}

	for _, brace := range braces {
		_, leftBrace := bracesPairs[brace]

		if leftBrace {
			stackBraces.Push(brace)
		} else {
			if stackBraces.Empty() {
				return false
			}
			if pairBrace := stackBraces.Pop().(string); bracesPairs[pairBrace] != brace {
				return false
			}
		}
	}
	return stackBraces.Empty()
}


func main()  {
	fmt.Println(valid("((){}[]){[]()[]}"))
	fmt.Println(valid("}((){}[]){[]"))
}
