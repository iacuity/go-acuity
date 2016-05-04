package util

type IBinarySearch interface {
	// Len is the number of elements in the collection.
	Len() int

	// Compare two object and return 0 slice[i] == element
	// return 1 if slice[i] > element
	// return -1 if slice[i] < element
	Compare(i int, element interface{}) int
}

// Generic Binary Search Implementation
func BinarySearch(slice IBinarySearch, element interface{}) int {
	length := slice.Len()

	if length < 0 {
		return -1
	}

	left, right := 0, (length - 1)

	for left <= right {
		mid := (left + right) / 2
		ret := slice.Compare(mid, element)
		switch ret {
		case 0:
			return mid
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		}
	}

	return -1
}
