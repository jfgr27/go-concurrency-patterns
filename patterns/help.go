package main

import (
	"log"
	"time"
)

func timeit(name string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Printf("%s took %s", name, duration)
	}
}
