# goFun
some funny tries with golang

[![Go Report Card](https://goreportcard.com/badge/github.com/fighterlyt/goFun)](https://goreportcard.com/report/github.com/fighterlyt/goFun)
[![codecov](https://codecov.io/gh/fighterlyt/goFun/branch/master/graph/badge.svg)](https://codecov.io/gh/fighterlyt/goFun)
[![GoDoc](https://godoc.org/github.com/fighterlyt/goFun?status.svg)](https://godoc.org/github.com/fighterlyt/goFun)


## Mutex with Timer
When we use **sync.Mutex**, we want to count the time cost by **Mutex.Lock** operation. When application comes to production,we can turn off the timing behavior.

This Mutex implementation has server properties:

*   you can turn off/on the timing behavior by **mutex.SetDebug()**
*   when turned off , it just behaves like **sync.Mutex**
*   when turned on,  it will record the time cost by **mutex.Mutex.Lock()***
    *   you can get the time duration by invoke **mutex.Mutex.Waited()**
    *   you can reset the timing by invoke **mutex.Mutex.Reset()** 
*   the precision used by mutex is milliSecond
