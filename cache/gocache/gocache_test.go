package gocache

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkSet(b *testing.B) {
	NewGoCache(5*time.Minute, 10*time.Minute)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		Set(key, value, 2*time.Minute)
	}
}

func BenchmarkGet(b *testing.B) {
	NewGoCache(5*time.Minute, 10*time.Minute)

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		Set(key, value, 2*time.Minute)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		Get(key)
	}
}

func BenchmarkDelete(b *testing.B) {
	NewGoCache(5*time.Minute, 10*time.Minute)

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		Set(key, value, 2*time.Minute)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		Delete(key)
	}
}

// Benchmark results
// goos: linux
// goarch: amd64
// pkg: github.com/rahulshewale153/common-lib/cache/gocache
// cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
// BenchmarkSet-8           2046751               662.8 ns/op
// BenchmarkGet-8           4092301               293.2 ns/op
// BenchmarkDelete-8        4407025               266.9 ns/op
// PASS
// ok      github.com/rahulshewale153/common-lib/cache/gocache   14.383s
