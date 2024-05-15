package redis

import "data_dump/ratelimit"

type RedisTokenBucket struct {
	ratelimit.TokenBucket

	key int
}

func (r *RedisTokenBucket) Acquire(leak_rate int) bool {
	return r.TokenBucket.Acquire(leak_rate)
}
