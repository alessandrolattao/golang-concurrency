package main

import (
	"fmt"
	"time"
)

func process(bestChannelInTheWorld chan string) {
	for task := range bestChannelInTheWorld {
		time.Sleep(10 * time.Second)
		fmt.Printf("Processed task: %s\n", task)
	}
}

func generate(bestChannelInTheWorld chan string, name string) {
	for {
		microTimestamp := time.Now().Format("2006-01-02 15:04:05.000000")
		taskName := fmt.Sprintf("%s --> %s", name, microTimestamp)
		select {
		case bestChannelInTheWorld <- taskName:
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func main() {
	bestChannelInTheWorld := make(chan string, 1)

	go process(bestChannelInTheWorld)

	go generate(bestChannelInTheWorld, "first")

	go generate(bestChannelInTheWorld, "second")

	select {}
}
