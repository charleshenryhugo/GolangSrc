package main

import (
	"bytes"
	"sync"
)

type info struct {
	Rwmu sync.RWMutex
	mu   sync.Mutex
	name string
}

type syncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func update(pInfo *info) {
	pInfo.mu.Lock()
	pInfo.name = "new name1"
	pInfo.mu.Unlock()
}
