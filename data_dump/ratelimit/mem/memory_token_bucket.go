package mem

import (
	"data_dump/ratelimit"
	"fmt"
)

// ToDo: fix some naming conventions
type TokenBucketMap interface {
	Append(key int, bucket *ratelimit.TokenBucket)
	Pop(key int) ratelimit.TokenBucket
	Acquire(key int, leakRate int) (bool, error)
}

type MemoryTokenBucket struct {
	tokens map[int]*ratelimit.TokenBucket
}

func newBucket() *MemoryTokenBucket {
	tokenMap := make(map[int]*ratelimit.TokenBucket)
	return &MemoryTokenBucket{tokens: tokenMap}
}

func (m *MemoryTokenBucket) Acquire(key int, leakRate int) (bool, error) {
	bucketValue := m.tokens[key]

	if bucketValue == nil {
		return false, fmt.Errorf("bucket does not exist")
	}

	return bucketValue.Acquire(leakRate), nil
}
