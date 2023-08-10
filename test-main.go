package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := 0

	fmt.Println(reflect.TypeOf(v).String())
}
