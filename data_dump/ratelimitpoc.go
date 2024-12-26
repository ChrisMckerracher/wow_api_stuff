package main

import (
	"data_dump/ratelimit"
	"data_dump/ratelimit/mem"
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
	"runtime"
)

var tokenBucket mem.TokenBucketMap
var fixedKey = 1

func sayHello() func() {
	threadVal := rand.IntN(10000)
	return func() {
		if success, err := tokenBucket.Acquire(fixedKey, 1); success && err == nil {
			fmt.Println(strconv.Itoa(threadVal))
			runtime.Gosched()
		} else {
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("Starting")
	keyBucket := ratelimit.TokenBucket{
		BucketSize: 10.0,
		BucketFill: 0.0,
		FillPeriod: 1000,
		LastUpdate: 0,
		FillRate:   1,
		Lock:       sync.Mutex{},
	}
	tokenBucket = mem.NewBucket()
	tokenBucket.Append(fixedKey, &keyBucket)

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			newHello := sayHello()

			for {
				newHello()
			}
		}()
	}

	select {}
}
