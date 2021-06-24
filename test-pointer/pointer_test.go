package testpointer

import (
	"fmt"
	"testing"
)

type MapT = map[string]int

func Test1(t *testing.T) {
	a := 111

	testInt(&a)
	fmt.Println("outside func")
	fmt.Println(a)

	b := MapT{"A": 1, "B": 2}
	testMap(b)
	fmt.Println("outside func")
	fmt.Println(b)

	c := Foo{id: 1, name: "foo"}
	d := new(Foo)
	d.id = 1
	d.name = "foo"
	testFoo(c, d)
	fmt.Println("outside func")
	fmt.Println(c, d)
}

func testInt(a *int) {
	fmt.Println("inside func")
	*a = 222
	fmt.Println(*a)
}

//https://stackoverflow.com/questions/40680981/are-maps-passed-by-value-or-by-reference-in-go
// You don't need to use a pointer with a map.
// Map types are reference types, like pointers or slices[1] ...
// Golang: Accessing a map using its reference, https://stackoverflow.com/questions/28384343/golang-accessing-a-map-using-its-reference
func testMap(b MapT) {
	fmt.Println("inside func")
	b["C"] = 33
	fmt.Println(b)
}

type Foo struct {
	id   int
	name string
}

func testFoo(c Foo, d *Foo) {
	fmt.Println("inside func")
	c.name = "bar"
	d.name = "bar"
	fmt.Println(c, *d)
}
