package main

import (
	"fmt"
	"reflect"
)

/*
初识反射
 */
func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

type Bird struct {
	Name string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}