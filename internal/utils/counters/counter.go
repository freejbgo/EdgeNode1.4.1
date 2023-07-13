// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package counters

import (
	syncutils "github.com/TeaOSLab/EdgeNode/internal/utils/sync"
	"github.com/cespare/xxhash"
	"github.com/iwind/TeaGo/Tea"
	"runtime"
	"time"
)

type Counter struct {
	countMaps uint64
	locker    *syncutils.RWMutex
	itemMaps  []map[uint64]*Item

	gcTicker *time.Ticker
	gcIndex  int
}

// NewCounter create new counter
func NewCounter() *Counter {
	var count = runtime.NumCPU() * 4
	if count < 8 {
		count = 8
	}

	var itemMaps = []map[uint64]*Item{}
	for i := 0; i < count; i++ {
		itemMaps = append(itemMaps, map[uint64]*Item{})
	}

	var counter = &Counter{
		countMaps: uint64(count),
		locker:    syncutils.NewRWMutex(count),
		itemMaps:  itemMaps,
	}

	return counter
}

// WithGC start the counter with gc automatically
func (this *Counter) WithGC() *Counter {
	if this.gcTicker != nil {
		return this
	}
	this.gcTicker = time.NewTicker(10 * time.Second)
	if Tea.IsTesting() {
		this.gcTicker = time.NewTicker(1 * time.Second)
	}
	go func() {
		for range this.gcTicker.C {
			this.GC()
		}
	}()

	return this
}

// Increase key
func (this *Counter) Increase(key uint64, lifeSeconds int) uint64 {
	var index = int(key % this.countMaps)
	this.locker.RLock(index)
	var item = this.itemMaps[index][key]
	this.locker.RUnlock(index)
	if item == nil { // no need to care about duplication
		item = NewItem(lifeSeconds)
		this.locker.Lock(index)

		// check again
		oldItem, ok := this.itemMaps[index][key]
		if !ok {
			this.itemMaps[index][key] = item
		} else {
			item = oldItem
		}

		this.locker.Unlock(index)
	}

	this.locker.Lock(index)
	var result = item.Increase()
	this.locker.Unlock(index)
	return result
}

// IncreaseKey increase string key
func (this *Counter) IncreaseKey(key string, lifeSeconds int) uint64 {
	return this.Increase(this.hash(key), lifeSeconds)
}

// Get value of key
func (this *Counter) Get(key uint64) uint64 {
	var index = int(key % this.countMaps)
	this.locker.RLock(index)
	defer this.locker.RUnlock(index)
	var item = this.itemMaps[index][key]
	if item != nil {
		return item.Sum()
	}
	return 0
}

// GetKey get value of string key
func (this *Counter) GetKey(key string) uint64 {
	return this.Get(this.hash(key))
}

// Reset key
func (this *Counter) Reset(key uint64) {
	var index = int(key % this.countMaps)
	this.locker.RLock(index)
	var item = this.itemMaps[index][key]
	this.locker.RUnlock(index)

	if item != nil {
		this.locker.Lock(index)
		delete(this.itemMaps[index], key)
		this.locker.Unlock(index)
	}
}

// ResetKey string key
func (this *Counter) ResetKey(key string) {
	this.Reset(this.hash(key))
}

// TotalItems get items count
func (this *Counter) TotalItems() int {
	var total = 0

	for i := 0; i < int(this.countMaps); i++ {
		this.locker.RLock(i)
		total += len(this.itemMaps[i])
		this.locker.RUnlock(i)
	}

	return total
}

// GC garbage expired items
func (this *Counter) GC() {
	var gcIndex = this.gcIndex

	this.gcIndex++
	if this.gcIndex >= int(this.countMaps) {
		this.gcIndex = 0
	}

	this.locker.RLock(gcIndex)
	var itemMap = this.itemMaps[gcIndex]
	var expiredKeys = []uint64{}
	for key, item := range itemMap {
		if item.IsExpired() {
			expiredKeys = append(expiredKeys, key)
		}
	}
	this.locker.RUnlock(gcIndex)

	if len(expiredKeys) > 0 {
		for _, key := range expiredKeys {
			this.locker.Lock(gcIndex)
			delete(itemMap, key)
			this.locker.Unlock(gcIndex)
		}
	}
}

// calculate hash of the key
func (this *Counter) hash(key string) uint64 {
	return xxhash.Sum64String(key)
}
