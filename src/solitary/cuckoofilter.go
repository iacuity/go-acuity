package solitary

import (
	"github.com/seiflotfy/cuckoofilter"
)

type CuckooContainer struct {
	container *cuckoofilter.CuckooFilter
}

// Return new container with give capacity
func newCuckooContainer(capacity uint) *CuckooContainer {
	return &CuckooContainer{
		container: cuckoofilter.NewCuckooFilter(capacity),
	}
}

// Return the size of container
func (cn *CuckooContainer) Size() uint {
	return cn.container.Count()
}

func (cn *CuckooContainer) Insert(key string) bool {
	return cn.container.InsertUnique([]byte(key))
}

func (cn *CuckooContainer) Delete(key string) bool {
	return cn.container.Delete([]byte(key))
}

func (cn *CuckooContainer) Exists(key string) bool {
	return cn.container.Lookup([]byte(key))
}
