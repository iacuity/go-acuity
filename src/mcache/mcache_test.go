package mcache_test

import (
	"mcache"
	"testing"

	"time"
)

func TestCache(t *testing.T) {
	defaultExpiry := 60 // in second
	err := mcache.Init(defaultExpiry)

	if nil != err {
		t.Fatal("Failed to Initialize Cache", err)
	}

	key := "foo"
	val := "baar"

	cval, ok := mcache.Get(key)

	if ok {
		t.Fatal("Garbage value for cached key")
	}

	mcache.Set(key, val)

	cval, ok = mcache.Get(key)

	if ok && val != cval {
		t.Fatal("Invalid value for cached key")
	}

	if !ok && val == cval {
		t.Fatal("Invalid value for cached key")
	}

	mcache.SetWithExpiry(key, val, 2*time.Second)

	cval, ok = mcache.Get(key)

	if ok && val != cval {
		t.Fatal("Invalid value for cached key")
	}

	if !ok && val == cval {
		t.Fatal("Invalid value for cached key")
	}

	time.Sleep(3 * time.Second)

	cval, ok = mcache.Get(key)

	if ok {
		t.Fatal("Garbage value for cached key")
	}
}
