package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/hurtki/teechan/teechan"
)

func main() {
	// ch - some channel ( chan int ) 
	ch := generator()
	// our goal is to create from it two channels
	// we are using a struct in /teechan/teechan.go
	tch := teechan.NewTeeChan[int](2)
	// giving to teechan a channel we need to "multiply"
	chs := tch.Execute(ch)
	// now we got two channels with the same outs as for first channel
	
	

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// now we are emulating to more gorutines that needed to read from the first channel 
	// so now first gorutine uses chs[0]
	// and second gorutine uses chs[1]
	
	go func() {
		defer wg.Done()
		for i := range chs[0] {
			time.Sleep(time.Second)
			fmt.Println("testing", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := range chs[1] {
			fmt.Println("monitoring", i)
		}
	}()
	
	wg.Wait()
}


// generator() returns a channel and writes into three times numbers (0, 1, 2)
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
