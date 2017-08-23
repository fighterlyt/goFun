package batchWriter

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	data = []byte(`abcdefghijklmnopqrstuvwxyz`)
)

func TestBatchWriter_Write(t *testing.T) {
	buffer := &bytes.Buffer{}

	w := NewBatchWriter(buffer, 50, time.Second*100)

	count, err := w.Write(data)

	assert.NoError(t, err, "写入 1/26/0/50")
	assert.Equal(t, count, len(data), "写入 1/26/0/50")

	assert.Equal(t, buffer.Len(), 0, "写入 1/26/0/50")

	count, err = w.Write(data)

	assert.NoError(t, err, "写入 2/26/26/50")
	assert.Equal(t, count, len(data), "写入 2/26/26/50")

	assert.Equal(t, buffer.Len(), 52, "写入 2/26/26/50")
}
func TestBatchWriter_Write2(t *testing.T) {
	SetDebug(true)

	buffer := &bytes.Buffer{}

	w := NewBatchWriter(buffer, 50, time.Second*1)

	count, err := w.Write(data)

	assert.NoError(t, err, "写入 1/26/0/50")
	assert.Equal(t, count, len(data), "写入 1/26/0/50")

	assert.Equal(t, buffer.Len(), 0, "写入 1/26/0/50")

	time.Sleep(time.Millisecond * 4100)

	assert.Equal(t, 26, buffer.Len(), "写入 2/26/26/50")
}

func TestMain(m *testing.M) {
	m.Run()
}

var (
	datas = map[string][]interface{}{
		"默认": {false, "默认debug关闭",emptyLog, log, "默认日志"},
		"打开": {true, "打开后",simpleLog, log, "默认日志"},
	}
)

func TestSetDebug(t *testing.T) {
	//默认状态
	infos := datas["默认"]

	assert.Equal(t, infos[0], debug, infos[1])
	//assert.Equal(t, infos[3], infos[4], infos[5])

	//打开
	SetDebug(true)
	infos = datas["打开"]

	assert.Equal(t, infos[0], debug, infos[1])

	//assert.Equal(t, infos[3], infos[4], infos[5])

	//再次关闭
	SetDebug(false)
	infos = datas["默认"]

	assert.Equal(t, infos[0], debug, infos[1])
}
