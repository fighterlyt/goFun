package mutex

import (
	"sync"
	"testing"
	"time"
)

var (
	result = 1
)

func TestMutex_Waited(t *testing.T) {
	SetDebug(true)
	m := &Mutex{}

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			m.Lock()

			time.Sleep(time.Millisecond * 100)

			m.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	t.Log(m.Waited().String(), m.waited)
}

func BenchmarkMutex_Lock(b *testing.B) {
	b.Run("sync.Mutex", func(b *testing.B) {
		wg := &sync.WaitGroup{}
		m := &sync.Mutex{}
		for i := 0; i < 20; i++ {
			go func() {
				wg.Add(1)
				m.Lock()

				time.Sleep(time.Millisecond * 100)
				result = 2
				m.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	})
	b.Run("Mutex", func(b *testing.B) {
		wg := &sync.WaitGroup{}
		m := &Mutex{}
		for i := 0; i < 20; i++ {
			go func() {
				wg.Add(1)
				m.Lock()

				time.Sleep(time.Millisecond * 100)
				result = 2

				m.Unlock()
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
