package teechan

import "sync"

type TeeChan[T any] struct {
	OutChansCount int
	chs           []chan T
}

func NewTeeChan[T any](outs int) *TeeChan[T] {
	chs := make([]chan T, outs)
	for i := 0; i < outs; i++ {
		chs[i] = make(chan T)
	}


	return &TeeChan[T]{
		OutChansCount: outs,
		chs:           chs,
	}
}

func (t *TeeChan[T]) Execute(ch chan T) []chan T {

	go func() {
		wg := &sync.WaitGroup{}
		for item := range ch {
			for i := range t.chs {
				wg.Add(1)
				go func(ch chan T, val T) {
					defer wg.Done()
					ch <- val
				}(t.chs[i], item)
			}
		}

		wg.Wait()

		for _, ch := range t.chs {
			close(ch)
		}

	}()

	return t.chs
}
