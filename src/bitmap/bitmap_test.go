package bitmap_test

import (
	"bitmap"
	"testing"
)

func TestConfig(t *testing.T) {
	pageSize := 4 * 1024            // 4kb page size
	volumnSize := 100 * 1024 * 1024 // 100 GB

	totalPages := volumnSize / pageSize

	t.Log("Total Pages: ", totalPages)

	bitMap := bitmap.NewBitMap(totalPages)
	testBitMap(t, bitMap)

	optimizedBitMap := bitmap.NewOptimizedBitMap(totalPages)
	testBitMap(t, optimizedBitMap)
}

func checkForError(t *testing.T, err error) {
	if nil != err {
		t.Fatal(err)
	}
}

func testBitMap(t *testing.T, bitMap bitmap.IBitMap) {
	var err error
	// page 50 is updated, should set bit in index 0, offset 50
	err = bitMap.SetBit(50)
	checkForError(t, err)

	// page 128 is updated, should set bit in index 2, offset 0
	err = bitMap.SetBit(128)
	checkForError(t, err)

	// page 156 is updated, should set bit in index 2, offset 28
	err = bitMap.SetBit(156)
	checkForError(t, err)

	// page 1023 is updated, should set bit in index 15, offset 63
	err = bitMap.SetBit(1023)
	checkForError(t, err)

	// page 8195 is updated, should set bit in index 128, offset 3
	err = bitMap.SetBit(8195)
	checkForError(t, err)

	// page 25000 is updated, should set bit in index 390, offset 40
	err = bitMap.SetBit(25000)
	checkForError(t, err)

	err = bitMap.SetBit(25600)
	if nil == err {
		t.Fatal("BitMap Index out of bound condition not handled")
	}

	//t.Log(bitMap.PrintBitMap())

	snapshot := bitMap.Snapshot()

	for key, value := range snapshot {
		t.Logf("PageIndex: %10d Value: %064b", key, value)
	}

	bitMap.Reset()
}
