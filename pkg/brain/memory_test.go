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

func TestBasic(t *testing.T) {
	manager := setupMemoryManager(t)

	if manager != GetMemoryManager() {
		t.Errorf("manager should be singleton")
	}

	key, err := manager.Remember("foobar")
	if err != nil {
		t.Errorf("cannot remember")
	}

	thing, err := manager.GetFromId(key)
	if err != nil {
		t.Errorf("cannot get from key %s", key)
	}
	if thing.(string) != "foobar" {
		t.Errorf("get: expected foobar got %q", thing)
	}

	err = manager.Forget(key)
	if err != nil {
		t.Errorf("cannot forget key %s", key)
	}
	_, err = manager.GetFromId(key)
	if err == nil {
		t.Errorf("cannot forget %s: %q", key, err)
	}
}

func TestRememberWithName(t *testing.T) {
	manager := setupMemoryManager(t)

	var (
		set       bool
		err       error
		retrieved interface{}
	)

	name := "test_key"
	thing := "test_thing"
	thingOverwrite := "test_thing_overwrite"

	set, err = manager.RememberWithName(name, thing, false)
	if !set || err != nil {
		t.Errorf("cannot remember with name")
	}

	set, err = manager.RememberWithName(name, thing, false)
	if set {
		t.Errorf("it should not overwrite the memory")
	}
	if err != nil {
		t.Fatal(err)
	}

	set, err = manager.RememberWithName(name, thingOverwrite, true)
	if !set || err != nil {
		t.Errorf("cannot overwrite it")
	}
	retrieved, err = manager.GetFromId(name)
	if err != nil {
		t.Fatal(err)
	}
	if retrieved.(string) != thingOverwrite {
		t.Errorf("get: expected %s got %q", thingOverwrite, retrieved)
	}

	err = manager.Forget(name)
	if err != nil {
		t.Errorf("cannot forget key %s", name)
		t.Fatal(err)
	}

}
