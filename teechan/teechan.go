package teechan

import "sync"

type TeeChan[T any] struct {
	OutChansCount int
	chs           []chan T
	wgs           []*sync.WaitGroup
}

func NewTeeChan[T any](outs int) *TeeChan[T] {
	chs := make([]chan T, outs)
	for i := 0; i < outs; i++ {
		chs[i] = make(chan T)
	}
	wgs := make([]*sync.WaitGroup, outs)
	for i := 0; i < outs; i++ {
		wgs[i] = &sync.WaitGroup{}
	}

	return &TeeChan[T]{
		OutChansCount: outs,
		chs:           chs,
		wgs:           wgs,
	}
}

func (t *TeeChan[T]) Execute(ch chan T) []chan T {

	go func() {

		for item := range ch {
			for i := 0; i < t.OutChansCount; i++ {
				t.wgs[i].Add(1)
				go func(ch chan T, val T, i int) {
					defer t.wgs[i].Done()
					ch <- val
				}(t.chs[i], item, i)
			}
		}

		for i := 0; i < t.OutChansCount; i++ {
			go func(i int) {
				t.wgs[i].Wait()
				close(t.chs[i])
			}(i)
		}

	}()

	return t.chs
}
