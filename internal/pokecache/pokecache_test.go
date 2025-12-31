package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second

	cache := NewCache(interval)

	key := "https://example.com"
	val := []byte("testdata")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected to find key %q in cache", key)
	}

	if string(got) != string(val) {
		t.Fatalf("expected value %q, got %q", string(val), string(got))
	}
}
