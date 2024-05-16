package ratelimit

import (
	"sync"
	"time"
)

type TokenBucketInterface interface {
	Acquire(leakRate int) bool
}

type TokenBucket struct {
	bucketSize float32
	bucketFill float32
	fillPeriod int64 // fill period in millis
	lastUpdate int64 // in epoch
	fillRate   int64 // :(
	lock       sync.Mutex
}

func (t *TokenBucket) Acquire(leakRate int) bool {
	leakRateFloat := float32(leakRate)
	t.lock.Lock()
	defer t.lock.Unlock()

	now := time.Now().UnixMilli()
	delta := now - t.lastUpdate
	t.bucketFill += min(float32((delta/t.fillPeriod)*t.fillRate), t.bucketSize)

	if t.bucketFill >= leakRateFloat {
		t.bucketFill -= leakRateFloat
		return true
	}

	return false
}
