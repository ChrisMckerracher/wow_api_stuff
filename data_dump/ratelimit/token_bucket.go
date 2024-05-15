package ratelimit

import "sync"

type TokenBucketInterface interface {
	Acquire(leak_rate int) bool
}

type TokenBucket struct {
	bucket_size int
	bucket_fill int
	fill_period int // fill period in millis
	fill_rate   int
	lock        sync.Mutex
}

func (t *TokenBucket) Acquire(leak_rate int) bool {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.bucket_fill += t.fill_rate * t.fill_period // ToDo: fix this line

	if t.bucket_fill > 0 {
		t.bucket_fill -= leak_rate
		return true
	}

	return false
}
