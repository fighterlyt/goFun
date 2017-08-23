package batchWriter

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

var (
	debug = false
	log   = emptyLog
)

func SetDebug(d bool) {
	if d {
		log = simpleLog
	}else{
		log=simpleLog
	}
	debug=d
}

type Log func(...interface{})

func simpleLog(data ...interface{}) {
	fmt.Printf("%s", time.Now().Format("2006-01-02 15:04:05:006"))
	fmt.Println(data...)
}
func emptyLog(data ...interface{}) {

}

type batchWriter struct {
	writer   io.Writer
	buf      *bytes.Buffer
	duration time.Duration
	size     int
	*sync.RWMutex
}

func NewBatchWriter(writer io.Writer, size int, duration time.Duration) *batchWriter {
	result := &batchWriter{
		writer:   writer,
		buf:      bytes.NewBuffer(make([]byte, 0, size)),
		duration: duration,
		size:     size,
		RWMutex:  &sync.RWMutex{},
	}
	go func() {
		for {
			time.Sleep(result.duration)
			log("定时写入")
			result.RLock()

			result.write()
			result.RUnlock()
		}
	}()

	return result
}

func (b *batchWriter) Write(p []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	if count, err := b.buf.Write(p); err != nil {
		return count, err
	} else {
		if b.buf.Len() >= b.size {
			_,err= b.write()
			return count,err
		}
		return count, nil
	}
}

func (b *batchWriter) write() (int, error) {
	defer func(){
		b.buf.Reset()
	}()
	return b.writer.Write(b.buf.Bytes())
}
