package pool

import (
	"log"
)

type Pool struct {
	Size      int
	poolChan  chan interface{}
	NewObject func() (interface{}, error)
}

// Create object and save it into pool for future reuse
// Return error if fails to create pool
func (pool *Pool) Init() error {
	log.Println("Initializing pool")
	pool.poolChan = make(chan interface{}, pool.Size)

	for i := 0; i < pool.Size; i++ {
		object, err := pool.NewObject()
		if nil != err {
			return err
		}

		pool.poolChan <- object
	}

	log.Println("Successfully initialized pool.")
	return nil
}

// Return object from pool
func (pool *Pool) Borrow() interface{} {
	return <-pool.poolChan
}

// Release object to pool
func (pool *Pool) Release(object interface{}) {
	pool.poolChan <- object
}
