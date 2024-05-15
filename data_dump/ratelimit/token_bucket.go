package ratelimit

import "sync"

type TokenBucketInterface interface {
	Acquire(leakRate int) bool
}

type TokenBucket struct {
	bucketSize int
	bucketFill int
	fillPeriod int // fill period in millis
	fillRate   int
	lock       sync.Mutex
}

func (t *TokenBucket) Acquire(leakRate int) bool {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.bucketFill += t.fillRate * t.fillPeriod // ToDo: fix this line

	if t.bucketFill > 0 {
		t.bucketFill -= leakRate
		return true
	}

	return false
}
