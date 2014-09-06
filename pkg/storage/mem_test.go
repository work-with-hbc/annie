/*
In memory storage support testcases.
*/

package storage

import (
	"testing"
)

func TestMemSimple(t *testing.T) {
	storage := makeMemStorage()
	defer storage.Terminate()

	testStorage("mem", storage, t)
}
