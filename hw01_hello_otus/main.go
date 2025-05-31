package main

import "fmt"

func reverse(s string) string {
	r := make([]byte, len(s))
	for i := range len(s) {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverse("Hello, OTUS!"))
}
