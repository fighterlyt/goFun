package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	debug = false
)

// SetDebug  make Mutex timing for Lock
func SetDebug(d bool) {
	debug = d
}

// Mutex which extends sync.Mutex with the ability to time the Lock operation
type Mutex struct {
	sync.Mutex
	waited int64 //等待时间，毫秒
}

// Lock fetch the Lock, if debug is true, sum the costing time
func (m *Mutex) Lock() {
	if debug {
		start := time.Now()
		m.Mutex.Lock()
		costMilliSeconds := time.Since(start).Nanoseconds() / (1000 * 1000)
		atomic.AddInt64(&m.waited, costMilliSeconds)
	} else {
		m.Mutex.Lock()
	}
}

// Unlock just unlock
func (m *Mutex) Unlock() {
	m.Mutex.Unlock()
}

//Waited get the waited time cost by Lock()
func (m *Mutex) Waited() time.Duration {
	d, _ := time.ParseDuration(fmt.Sprintf("%dms", atomic.LoadInt64(&m.waited)))
	return d
}

//Reset reset the timer
func (m *Mutex) Reset() {
	atomic.StoreInt64(&m.waited, 0)
}
