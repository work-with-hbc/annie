/*
Storage testcases
*/

package storage

import "testing"

func testStorage(name string, storage Storage, t *testing.T) {
	if storage.Save("test", []byte("foobar")) != nil {
		t.Errorf("%s cannot save", name)
	}

	thing, err := storage.Get("test")
	if err != nil {
		t.Errorf("%s cannot get", name)
	}
	if string(thing) != "foobar" {
		t.Errorf("%s get: expected: foobar, got %s", name, string(thing))
	}

	if storage.Forget("test") != nil {
		t.Errorf("%s cannot forget", name)
	}
	thing, err = storage.Get("test")
	if err == nil {
		t.Errorf("%s didn't forget, got %s", name, string(thing))
	}
}
