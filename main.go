package main

import (
	"time"

	"log"
)

func main() {
	index := 0
	for {
		time.Sleep(5 * time.Second)
		index++
		log.Printf("hello world! Pass: %d", index)
	}

}
