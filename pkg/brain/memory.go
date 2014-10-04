/*
# Memory support.

Annie brain memory is a k-v database.

## Data types

Currently annie supports store things with these types:

- String.
- Lists: collections of strings.

## Data type operations

- list: list everything from the storage.
- has key: check if the storage contains an item with the key.
- forget: forget an item by given key.

### String

- store: store a string and retrieve a uuid key.
- store by key: store a string by given key. It will overwrite the existed item with that key.
- get by key.

### Lists

- store: store a list and retrieve a uuid key.
- store by key: store a list by given key. It will overwrite the existed item with that key.
- push: push a string into specific list.
- get by key.
*/

package brain

import (
	"encoding/json"
	"fmt"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/work-with-hbc/annie/pkg/jsonconfig"
	"github.com/work-with-hbc/annie/pkg/storage"
)

type MemoryManager struct {
	storage storage.Storage
}

func (m MemoryManager) getFromStorageByKey(key string) ([]byte, error) {
	return m.storage.Get(key)
}

func (m MemoryManager) writeToStorageByKey(key string, thing []byte) error {
	return m.storage.Save(key, thing)
}

func (m MemoryManager) generateUniqueKey() (key string) {
	for {
		key = uuid.New()
		if !m.Has(key) {
			return key
		}
	}

	return "unable-to-generate-unique-key"
}

// TODO
func (m MemoryManager) List() ([]interface{}, error) {
	return nil, nil
}

func (m MemoryManager) Has(key string) bool {
	return m.storage.Has(key)
}

func (m MemoryManager) Forget(key string) error {
	if !m.Has(key) {
		return nil
	}

	return m.storage.Forget(key)
}

func (m MemoryManager) StoreStringByKey(key string, str string) error {
	return m.writeToStorageByKey(key, []byte(str))
}

func (m MemoryManager) StoreString(str string) (string, error) {
	key := m.generateUniqueKey()

	err := m.StoreStringByKey(key, str)
	if err != nil {
		key = ""
	}

	return key, err
}

func (m MemoryManager) GetString(key string) (string, error) {
	if !m.Has(key) {
		return "", nil
	}

	raw, err := m.getFromStorageByKey(key)
	return string(raw), err
}

func (m MemoryManager) StoreListByKey(key string, list []string) error {
	marshaledList, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return m.writeToStorageByKey(key, marshaledList)
}

func (m MemoryManager) StoreList(list []string) (string, error) {
	key := m.generateUniqueKey()

	err := m.StoreListByKey(key, list)
	if err != nil {
		key = ""
	}

	return key, err
}

func (m MemoryManager) PushList(key string, item string) error {
	if !m.Has(key) {
		return fmt.Errorf("Unable to get list from key: %s", key)
	}

	var list []string
	if err := m.GetList(key, &list); err != nil {
		return err
	}

	list = append(list, item)
	if err := m.StoreListByKey(key, list); err != nil {
		return err
	}

	return nil
}

func (m MemoryManager) GetList(key string, list *[]string) error {
	if !m.Has(key) {
		return nil
	}

	raw, err := m.getFromStorageByKey(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, list)
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
