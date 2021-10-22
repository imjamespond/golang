package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Foo struct {
	bar int
}

func (f *Foo) test() {
	fmt.Println(f.bar)
}

var (
	foo  *Foo
	once sync.Once
	wg   sync.WaitGroup
)

func instance() *Foo {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
		foo = &Foo{bar: rand.Int()}
	})
	return foo
}

func TestSingleton(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			inst := instance()
			inst.test()
			// fmt.Println(inst.Num)
			wg.Done()
		}()
	}
	wg.Wait()
}
