package pokecache

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
  createdAt time.Time
  value *json.Decoder 
}

type Cache struct {
  mu sync.Mutex
  values map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
  var cachedData Cache 
  cachedData.values = make(map[string]cacheEntry)

  go readLoop(interval, &cachedData)
  return &cachedData
}

func (cachedData *Cache) Add(key string, value *json.Decoder) {
  cachedData.mu.Lock()
  defer cachedData.mu.Unlock()

  cachedData.values[key] = cacheEntry{ createdAt: time.Now(), value: value }
}


func (cachedData *Cache) Get(key string) (*json.Decoder, error) {
  cachedData.mu.Lock()
  defer cachedData.mu.Unlock()

  val, ok := cachedData.values[key]
  if !ok {
    return nil, errors.New("Not found in cache")
  }

  return val.value, nil
}

func readLoop(interval time.Duration, cachedData *Cache) {
  ticker := time.NewTicker(5 * time.Second)
  for {
    for key, value := range cachedData.values {
      if value.createdAt.Before(time.Now().Add(-interval)) {
        cachedData.mu.Lock()
        delete(cachedData.values, key) 
        cachedData.mu.Unlock()
      }
    }

    <-ticker.C  
  }
}
