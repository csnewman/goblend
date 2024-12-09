package main

import (
	"log"

	"github.com/csnewman/goblend"
)

func main() {
	log.Println("Add example")

	log.Println(goblend.Call(addPtr, 0))

	log.Println("End of example")
}
