package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now().UnixNano())
			time.Sleep(time.Millisecond * 500)
		}
	}
}
