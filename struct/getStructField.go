package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Bar string
	Baz int
}

func main() {
	list := []Foo{Foo{Bar: "hiahoma", Baz: 1}, Foo{Bar: "nizaiganma", Baz: 2}}

	var Foo_Bar = reflect.TypeOf(Foo{}).Field(0)
	// var Foo_Bar = reflect.TypeOf(Foo{}).Field(0).Name
	// var Foo_Bar = reflect.ValueOf(Foo{}).Type().Field(0).Name
	// var Foo_Baz = reflect.TypeOf(Foo{}).Field(1).Name
	fmt.Printf("%+v \n", Foo_Bar)

	for _, n := range list {
		v := reflect.ValueOf(n) //.Elem()
		// fmt.Printf("%+v \n", v.Type().Field(0).Name)
		// fmt.Printf("%+v \n", v.NumField())
		// fmt.Printf("%+v \n", v.Field(0))
		// fmt.Printf("%+v \n", fmt.Sprint(v.Field(0).Interface()))

		for i := 0; i < v.NumField(); i++ {
			fmt.Println(v.Field(i)) //.Name)
		}

	}
}
