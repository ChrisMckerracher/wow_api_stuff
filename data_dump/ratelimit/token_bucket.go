package ratelimit

import (
	"sync"
	"time"
)

type TokenBucketInterface interface {
	Acquire(leakRate int) bool
}

type TokenBucket struct {
	BucketSize float32
	BucketFill float32
	FillPeriod int64 // fill period in millis
	LastUpdate int64 // in epoch
	FillRate   int64 // :(
	Lock       sync.Mutex
}

func (t *TokenBucket) Acquire(leakRate int) bool {
	leakRateFloat := float32(leakRate)
	t.Lock.Lock()
	defer t.Lock.Unlock()

	now := time.Now().UnixMilli()
	delta := now - t.LastUpdate
	t.LastUpdate = now
	t.BucketFill += min(float32((delta/t.FillPeriod)*t.FillRate), t.BucketSize)

	if t.BucketFill >= leakRateFloat {
		t.BucketFill -= leakRateFloat
		return true
	}

	return false
}
