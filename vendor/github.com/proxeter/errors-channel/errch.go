package errch

import "github.com/chapsuk/wait"

func Register(fn func() error) <-chan error {
	ch := make(chan error)
	wg := wait.Group{}
	wg.Add(func() {
		ch <- fn()
	})

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
