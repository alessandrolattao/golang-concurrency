package main

import (
	"fmt"
	"time"
)

// process consumes the channel and processes the tasks
func process(bestChannelInTheWorld chan string) {
	for task := range bestChannelInTheWorld {
		time.Sleep(10 * time.Second)
		fmt.Printf("Processed: %s\n", task)
	}
}

// generate produces a timestamp every millisecond
func generate(bestChannelInTheWorld chan string) {
	for {
		microTimestamp := time.Now().Format("2006-01-02 15:04:05.000000")
		select {
		case bestChannelInTheWorld <- microTimestamp:
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	bestChannelInTheWorld := make(chan string, 1)

	go process(bestChannelInTheWorld)

	go generate(bestChannelInTheWorld)

	select {}
}
