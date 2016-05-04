package solitary_test

import (
	"testing"

	"solitary"
)

func TestContainer(t *testing.T) {
	container := solitary.NewContainer(100, solitary.CuckooFilterType)

	t.Log("Container Size is: ", container.Size())
	deleted := container.Delete("foo")
	if deleted {
		t.Fatal("Deleted non-existing key")
	}

	inserted := container.Insert("foo")

	if !inserted {
		t.Fatal("Not able to insert key")
	}

	t.Log("Container Size is: ", container.Size())

	found := container.Exists("foo")

	if !found {
		t.Fatal("Key not found even it exists")
	}

	deleted = container.Delete("foo")
	if !deleted {
		t.Fatal("Not able to delete key")
	}
	t.Log("Container Size is: ", container.Size())

	found = container.Exists("bar")

	if found {
		t.Fatal("Key found even it  not exists")
	}
}
