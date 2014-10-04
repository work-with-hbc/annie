/*
Testcases for memory.
*/

package brain

import (
	"testing"

	"github.com/work-with-hbc/annie/pkg/jsonconfig"
)

func setupMemoryManager(t *testing.T) *MemoryManager {
	config, err := jsonconfig.LoadFromString("{}")
	if err != nil {
		t.Fatal(err)
	}

	SetupMemoryManagerWithConfig(config)
	manager := GetMemoryManager()

	if manager == nil {
		t.Fatalf("cannot get memory manager")
	}

	return manager
}

func TestMananger(t *testing.T) {
	manager := setupMemoryManager(t)

	if manager != GetMemoryManager() {
		t.Errorf("manager should be singleton")
	}
}

// TODO
func TestList(t *testing.T) {}

func TestHas(t *testing.T) {
	manager := setupMemoryManager(t)

	key := "test"
	item := []byte("test")

	if manager.Has(key) {
		t.Errorf("should not have key: %s", key)
	}

	if err := manager.writeToStorageByKey(key, item); err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}
}

func TestForget(t *testing.T) {
	manager := setupMemoryManager(t)

	key := "test"
	item := []byte("test")

	if err := manager.writeToStorageByKey(key, item); err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}

	if err := manager.Forget(key); err != nil {
		t.Fatal(err)
	}

	if manager.Has(key) {
		t.Errorf("should not have key: %s", key)
	}
}

func TestStoreStringByKey(t *testing.T) {
	manager := setupMemoryManager(t)

	key := "test"
	item := "test"

	if err := manager.StoreStringByKey(key, item); err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}

	stored, err := manager.GetString(key)
	if err != nil {
		t.Fatal(err)
	}
	if stored != item {
		t.Errorf("expected: %s, got: %s", item, stored)
	}
}

func TestStoreString(t *testing.T) {
	manager := setupMemoryManager(t)

	item := "test"

	key, err := manager.StoreString(item)
	if err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}

	stored, err := manager.GetString(key)
	if err != nil {
		t.Fatal(err)
	}
	if stored != item {
		t.Errorf("expected: %s, got: %s", item, stored)
	}
}

func TestStoreListByKey(t *testing.T) {
	manager := setupMemoryManager(t)

	key := "test"
	items := []string{
		"test1",
		"test2",
	}

	if err := manager.StoreListByKey(key, items); err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}

	var stored []string
	err := manager.GetList(key, &stored)
	if err != nil {
		t.Fatal(err)
	}

	for i, item := range items {
		if item != stored[i] {
			t.Errorf("expected: %s, got: %s", item, stored)
		}
	}
}

func TestPushList(t *testing.T) {
	manager := setupMemoryManager(t)

	key := "test"
	items := []string{
		"test1",
		"test2",
	}
	pushed := "test3"

	if err := manager.StoreListByKey(key, items); err != nil {
		t.Fatal(err)
	}

	if !manager.Has(key) {
		t.Errorf("must have key: %s", key)
	}

	items = append(items, pushed)
	if err := manager.PushList(key, pushed); err != nil {
		t.Fatal(err)
	}

	var stored []string
	err := manager.GetList(key, &stored)
	if err != nil {
		t.Fatal(err)
	}

	for i, item := range items {
		if item != stored[i] {
			t.Errorf("expected: %s, got: %s", item, stored)
		}
	}
}
