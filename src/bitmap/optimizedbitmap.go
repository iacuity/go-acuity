package bitmap

import (
	"bytes"
	"fmt"
)

// TODO: we can implement our own Algorithm to the replacement of map
// Algorithm for map replacement
// 1. Use Array/Slice for bitMap storage e.g. make([]Pair, 10)
// 2. Implement Insert Method which will do sorted insertion for key
// 3. We can run Binary search to get right index for setting bit as we mainted sorted order
//    for bitMap storage
// 4. Array/Slice can be re-allocated when there is no space to add new element
// Where Pair structure will be as {Key, Value}. Key will be index and Value will uint64
type OptimizedBitMap struct {
	bitMap map[int]uint64
	size   int
}

// initialize empty OptimizedBitMap
func NewOptimizedBitMap(n int) *OptimizedBitMap {

	if n <= 0 || n > MAX_BITMAP_SIZE {
		n = MAX_BITMAP_SIZE
	}

	return &OptimizedBitMap{bitMap: make(map[int]uint64), size: n}
}

func (obMap *OptimizedBitMap) SetBit(bitIndex int) error {
	maxBitIndex := obMap.size - 1

	if bitIndex > maxBitIndex {
		return fmt.Errorf("Invalid bitIndex, bitIndex should not exceed %d", maxBitIndex)
	}

	index := bitIndex >> MAX_UNIT_SIZE // get the array index for setting bits
	offset := uint64(bitIndex - (index * UINT64_SIZE))

	obMap.bitMap[index] |= (1 << offset)

	return nil
}

// clear the OptimizedBitMap
func (obMap *OptimizedBitMap) Reset() {
	for index, _ := range obMap.bitMap {
		delete(obMap.bitMap, index)
	}
}

// return the current snapshot of updated pages
func (obMap *OptimizedBitMap) Snapshot() map[int]uint64 {
	return obMap.bitMap
}

// return string representation of BitMap for debugging
func (obMap *OptimizedBitMap) PrintBitMap() string {
	buff := bytes.Buffer{}

	for index, value := range obMap.bitMap {
		buff.WriteString(fmt.Sprintf("page-%10d - %064b\n", index, value))
	}

	return buff.String()
}
