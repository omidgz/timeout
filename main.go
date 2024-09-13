package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	done := 0
	for i := 1; i <= 100; i++ {
		go func() {
			wg.Add(1)
			if res, err := runWithTimeout(); err == nil {
				fmt.Printf("%s\t %d Done %s\n", time.Now().Format(time.UnixDate), done+1, res)
				done += 1 // needed this after print
			} else {
				fmt.Printf("%s\t %s\n", time.Now().Format(time.UnixDate), err)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func theFunc() string {
	r := rand.Intn(10)
	time.Sleep(time.Duration(r) * time.Second)
	return fmt.Sprintf("%d", r)
}

func runWithTimeout() (string, error) {
	done := make(chan bool)
	res := ""
	go func() {
		res = theFunc()
		done <- true
	}()

	select {
	case <-done:
		return res, nil
	case <-time.After(3 * time.Second):
		return "", fmt.Errorf("Timeout")
	}
}
