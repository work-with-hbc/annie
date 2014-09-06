/*
In memory storage support.
*/

package storage

import (
	"errors"

	"github.com/bcho/annie/pkg/jsonconfig"
)

type MemStorage struct {
	data map[string][]byte
}

func (m *MemStorage) Save(key string, thing []byte) error {
	m.data[key] = thing

	return nil
}

func (m *MemStorage) Get(key string) ([]byte, error) {
	value, present := m.data[key]
	if !present {
		return nil, errors.New("not found")
	}

	return value, nil
}

func (m *MemStorage) Has(key string) bool {
	_, present := m.data[key]
	return present
}

func (m *MemStorage) Forget(key string) error {
	delete(m.data, key)

	return nil
}

func (m *MemStorage) Terminate() error {
	return nil
}

func makeMemStorage() *MemStorage {
	m := new(MemStorage)
	m.data = make(map[string][]byte)

	return m
}

func MakeMemStorageWithConfig(config *jsonconfig.Config) Storage {
	return makeMemStorage()
}
