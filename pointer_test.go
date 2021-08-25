package main

import (
	"fmt"
	"testing"
)

func TestPassingArgs1(t *testing.T) {
	p := Person{Name: "hello", Age: 1}
	fmt.Println("Person", p)

	modifyPerson(p)
	fmt.Println("modifyPerson", p)

	modifyIPerson(&p)
	fmt.Println("modifyIPerson", p)

	modifyPersonByPt(&p)
	fmt.Println("modifyPersonByPt", p)
}

type IPerson interface {
	setName(name string)
	getName() string
}

type Person struct {
	Name string
	Age  int16
}

func (p *Person) setName(name string) {
	p.Name = name
}

func (p *Person) getName() string {
	return p.Name
}

func modifyPerson(p Person) { // pass new copy of Person
	p.Name = "bar"
	p.Age = 999
}

func modifyIPerson(p IPerson) {
	p.setName("far")
}

func modifyPersonByPt(p *Person) {
	p.setName("foobar")
}
