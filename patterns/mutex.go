package main


import (
	"fmt"
	"sync"
)
type MutexWork struct {
	m bool
}

type Container struct {
    mu       sync.Mutex
    counters map[string]int
}



func (c *Container) incMutex(name string) {
    c.mu.Lock() // thread not in lock will block here until free.
    defer c.mu.Unlock()
    c.counters[name]++
}

func (c *Container) incNoMutex(name string) {
    c.counters[name]++
}

func (m *MutexWork) work() {

	if m.m == true {
		m.mutexWork()
	} else {
		m.noMutexWork()
	}
}
func (m *MutexWork) mutexWork() {
	fmt.Sprintf("Mutex work")

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup


	doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.incMutex(name)
        }
        wg.Done()
	}

	wg.Add(3)


	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)


}


func (m *MutexWork) noMutexWork() {
	fmt.Sprintf("No Mutex work")

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}


	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.incNoMutex(name)
		}
		wg.Done()
	}


	wg.Add(3)

	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	fmt.Println(c.counters)



}
