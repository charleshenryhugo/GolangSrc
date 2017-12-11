package main

import (
	"bytes"
	"sync"
)

type Info struct {
	Rwmu sync.RWMutex
	mu   sync.Mutex
	name string
}

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func update(pInfo *Info) {
	pInfo.mu.Lock()
	pInfo.name = "new name1"
	pInfo.mu.Unlock()
}
