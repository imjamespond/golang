package testpointer

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	p := Person{"foo", 123}
	fmt.Println(p)
	modifyPerson(p)
	fmt.Println(p)
	modifyIPerson(p)
	fmt.Println(p)
	modifyPersonByPt(&p)
	fmt.Println(p)
}

type IPerson interface {
	setName(name string)
}

type Person struct {
	Name string
	Age  int16
}

func (p Person) setName(name string) {
	p.Name = name
}

func modifyPerson(p Person) {
	p.Name = "bar"
	p.Age = 999
}

func modifyIPerson(p IPerson) {
	p.setName("far")
}

func modifyPersonByPt(p *Person) {
	p.Name = "ear"
}
