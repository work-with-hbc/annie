/*
Memory support.
*/

package brain

import (
	"encoding/json"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/work-with-hbc/annie/pkg/jsonconfig"
	"github.com/work-with-hbc/annie/pkg/storage"
)

type MemoryManager struct {
	storage storage.Storage
}

func (m MemoryManager) getFromStorage(key string) (interface{}, error) {
	data, err := m.storage.Get(key)
	if err != nil {
		return nil, err
	}

	var value interface{}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (m MemoryManager) writeToStorage(key string, something interface{}) error {
	marshaled, err := json.Marshal(something)
	if err != nil {
		return err
	}

	return m.storage.Save(key, marshaled)
}

func (m MemoryManager) Remember(something interface{}) (string, error) {
	var itemId string
	for {
		itemId = uuid.New()
		if !m.storage.Has(itemId) {
			break
		}
	}

	err := m.writeToStorage(itemId, something)
	if err != nil {
		return "", err
	}

	return itemId, nil
}

// Write something to storage with name.
// Returns `false` if the name has been used and `overwrite` didn't set.
func (m MemoryManager) RememberWithName(name string, something interface{}, overwrite bool) (bool, error) {
	if m.storage.Has(name) && !overwrite {
		return false, nil
	}

	err := m.writeToStorage(name, something)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m MemoryManager) GetFromId(id string) (interface{}, error) {
	return m.getFromStorage(id)
}

func (m MemoryManager) Forget(id string) error {
	err := m.storage.Forget(id)
	if err != nil {
		return err
	}

	return nil
}

func makeMemoryManager() *MemoryManager {
	return &MemoryManager{}
}

// Singleton instance.
var instance *MemoryManager

func ensureHaveMemoryManaager() {
	if instance == nil {
		instance = makeMemoryManager()
	}
}

func SetupMemoryManagerWithConfig(config *jsonconfig.Config) *MemoryManager {
	ensureHaveMemoryManaager()

	storageMaker := storage.MakeMemStorageWithConfig

	switch strings.ToLower(config.GetString("storage.provider")) {
	case "leveldb":
		storageMaker = storage.MakeLevelDBStorageWithConfig
	}

	instance.storage = storageMaker(config.GetConfigObj("storage.options"))

	return instance
}

func GetMemoryManager() *MemoryManager {
	ensureHaveMemoryManaager()

	return instance
}
