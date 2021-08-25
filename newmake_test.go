package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var sp *string = new(string) //关键点
	*sp = "Golang"
	fmt.Println(*sp)
}

type person struct {
	name string
	age  int
}

func TestInitialization(t *testing.T) {
	{
		p := person{name: "张三", age: 18}
		fmt.Println(p)
	}
	{
		p := new(person)
		p.name = "李四"
		p.age = 20
		fmt.Println(*p)
	}
}

func TestMake(t *testing.T) {
	{
		ch := make(chan int)
		go func(chan int) {
			// time.Sleep(time.Second)
			ch <- 1
			ch <- 2
		}(ch)
		fmt.Println("Waiting...")
		fmt.Println(<-ch, <-ch)
	}
	{
		m := make(map[string]int, 2)
		m["foo"] = 1
		m["bar"] = 2
		fmt.Println(m)
	}
	{
		s := make([]int, 2, 3)
		s[0] = 1
		s[1] = 2
		fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
		s = append(s, 3)
		s = append(s, 4)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
	}
}
