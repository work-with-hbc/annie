/*
Leveldb storage support testcases.
*/

package storage

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLevelDBSimple(t *testing.T) {
	testPath, err := ioutil.TempDir("/tmp", "annie-test")
	if err != nil {
		t.Errorf("cannot create temp dir")
	}
	defer func() {
		err := os.RemoveAll(testPath)
		if err != nil {
			t.Errorf("cannot cleanup")
		}
	}()

	storage := makeLevelDBStorage(testPath)
	defer storage.Terminate()

	testStorage("leveldb", storage, t)
}
