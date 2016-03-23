package pool_test

import (
	"log"
	"pool"
	"sync"
	"testing"
	"time"
)

func NewString() (interface{}, error) {
	return "Tushar", nil
}

func worker(pool *pool.Pool, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	object := pool.Borrow()
	log.Println(i, "-", object)
	time.Sleep(1 * time.Millisecond)
	pool.Release(object)
}

func TestPool(t *testing.T) {
	pool := &pool.Pool{Size: 10,
		NewObject: NewString,
	}

	err := pool.Init()

	if nil != err {
		t.Fatal("Failed to create pool")
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go worker(pool, wg, i)
	}

	wg.Done()
	wg.Wait()
}
