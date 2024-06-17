package utils

import (
	"context"
	"sync"
)

// TTLMap is a thread-safe map to store data with context.
type TTLMap struct {
	data  map[string]interface{}
	mutex sync.Mutex
}

// NewDataStore initializes a new DataStore.
func NewDataStore() *TTLMap {
	return &TTLMap{
		data: make(map[string]interface{}),
	}
}

// Set adds a key-value pair to the DataStore with a context.
func (ds *TTLMap) Set(ctx context.Context, key string, value interface{}) {
	ds.mutex.Lock()
	ds.data[key] = value
	ds.mutex.Unlock()

	go func() {
		<-ctx.Done()
		ds.mutex.Lock()
		delete(ds.data, key)
		ds.mutex.Unlock()
	}()
}

// Get retrieves a value by key from the DataStore.
func (ds *TTLMap) Get(key string) (interface{}, bool) {
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	value, exists := ds.data[key]
	return value, exists
}
