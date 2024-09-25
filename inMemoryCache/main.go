package main

import (
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type inMemory struct {
	data map[string]string
	mu   sync.RWMutex
}

func New() Cache {
	return &inMemory{}
}

func (i *inMemory) Set(k, v string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[k] = v

}

func (i *inMemory) Get(k string) (v string, ok bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	val, ok := i.data[k]
	return val, ok
}
