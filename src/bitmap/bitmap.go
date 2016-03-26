package bitmap

import (
	"bytes"
	"fmt"
)

// Interface for BitMap implementation
type IBitMap interface {
	SetBit(index int) error
	Reset()
	Snapshot() map[int]uint64
	PrintBitMap() string
}

type BitMap struct {
	bits []uint64
	size int
}

const (
	MAX_UNIT_SIZE   = 6                  // 2^6 = 64
	UINT64_SIZE     = 1 << MAX_UNIT_SIZE // size of uint64
	MAX_BITMAP_SIZE = 1 << 24            // MAX VALUE WILL BE 2^32
)

// Allocates continuous n bits
// If value of n is 0 or negative then allocated BitMap for MAX_BITMAP_SIZE
func NewBitMap(n int) *BitMap {
	if n <= 0 || n > MAX_BITMAP_SIZE {
		n = MAX_BITMAP_SIZE
	}

	bitMapSize := n>>MAX_UNIT_SIZE + 1 // get reuired array size by divide by 2^6 = 64

	fmt.Println("Map size: ", bitMapSize)

	return &BitMap{bits: make([]uint64, bitMapSize),
		size: n,
	}

}

// set the bitIndex of BitMap
// return error if not able to set bitIndex of BitMap
func (bitMap *BitMap) SetBit(bitIndex int) error {
	maxBitIndex := bitMap.size - 1

	if bitIndex > maxBitIndex {
		return fmt.Errorf("Invalid bitIndex, bitIndex should not exceed %d", maxBitIndex)
	}

	index := bitIndex >> MAX_UNIT_SIZE // get the array index for setting bits
	offset := uint64(bitIndex - (index * UINT64_SIZE))

	bitMap.bits[index] |= (1 << offset)

	return nil
}

// reset all bits to zero
func (bitMap *BitMap) Reset() {
	arraySize := len(bitMap.bits)
	for index := 0; index < arraySize; index++ {
		bitMap.bits[index] = 0
	}
}

// return the current snapshot of updated pages
func (bitMap *BitMap) Snapshot() map[int]uint64 {
	snapshot := make(map[int]uint64)
	arraySize := len(bitMap.bits)

	for index := 0; index < arraySize; index++ {
		value := bitMap.bits[index]
		if 0 != value {
			snapshot[index] = bitMap.bits[index]
		}
	}

	return snapshot
}

// return string representation of BitMap for debugging
func (bitMap *BitMap) PrintBitMap() string {
	buff := bytes.Buffer{}

	arraySize := len(bitMap.bits)
	for index := 0; index < arraySize; index++ {
		buff.WriteString(fmt.Sprintf("page-%d - %064b\n", index, bitMap.bits[index]))
	}

	return buff.String()
}
