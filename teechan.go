package teechan



type TeeChan[T any] struct {
	OutChansCount int
	chs []chan T
}

func NewTeeChan[T any](outs int) *TeeChan[T] {
	chs := make([]chan T, outs)
	for i := range outs {
		chs[i] = make(chan T)
	}
	
	return &TeeChan[T]{
		OutChansCount: outs,
		chs: chs,
	}
}

func (t *TeeChan[T]) Execute(ch chan T) []chan T {
	for item := range ch {
		for _, ch := range t.chs {
			go func() {
				ch <- item
			}()
		}
	}
	return t.chs
}

