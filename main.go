package main

import (
    "fmt"
    "sync"
)

type Database struct {
    data map[string]string
    mu   sync.RWMutex
}

func NewDatabase() *Database {
    return &Database{
        data: make(map[string]string),
    }
}

func (db *Database) Get(key string) (string, bool) {
    db.mu.RLock()
    defer db.mu.RUnlock()
    value, exists := db.data[key]
    return value, exists
}

func main() {
    
    db := NewDatabase()

}