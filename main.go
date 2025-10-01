package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/hurtki/teechan/teechan"
)

func main() {
	tch := teechan.NewTeeChan[int](2)
	chs := tch.Execute(generator())
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range chs[0] {
			time.Sleep(time.Second * 2)
			fmt.Println("testing", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range chs[1] {
			time.Sleep(time.Second)
			fmt.Println("monitoring", i)
		}
	}()
	wg.Wait()

}

func generator() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
