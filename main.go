package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)

	for i := 1; i <= 50; i++ {
		go func() {
			r := rand.Intn(50)
			time.Sleep(time.Duration(r) * time.Second)
			c <- fmt.Sprintf("%d", r)
		}()
	}
	done := 0
	for {
		select {
		case res := <-c:
			done += 1
			fmt.Printf("%s\t%d done %s\n", time.Now().Format(time.UnixDate), done, res)
		case <-time.After(1 * time.Second):
			fmt.Printf("%s\tTimeout!\n", time.Now().Format(time.UnixDate))
		}
	}
}
