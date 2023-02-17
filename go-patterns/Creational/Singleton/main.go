package main

import "sync"

func main() {

	useMtx := false

	wg := &sync.WaitGroup{}
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			if useMtx {
				getInstanceWithMtx()
			} else {
				getInstanceWithOnce()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
