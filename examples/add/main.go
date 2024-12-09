package main

import (
	"log"
	"sync"

	"github.com/csnewman/goblend"
)

var wg = &sync.WaitGroup{}

func main() {
	log.Println("Add example")

	ptr := goblend.FuncPtr(ExampleCallback)

	log.Println(goblend.Call(addPtr, ptr))

	wg.Wait()

	log.Println("End of example")
}

func ExampleCallback(ptr *uint64) {
	log.Println("Hello world callback", *ptr)

	wg.Add(1)

	go func() {
		log.Println("Go routine was spawned")

		wg.Done()
	}()
}
