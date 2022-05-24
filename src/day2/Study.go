package main

import "fmt"

func main() {
	s := "2222"
	fmt.Println(s)
	u := s[2]
	fmt.Println(u)
	fmt.Println(u == byte('2'))
	fmt.Println(byte(2))
	fmt.Println(byte('2'))
}

func init() {
	fmt.Println("hello world")
}

