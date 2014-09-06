/*
Testcases for memory.
*/

package brain

import (
	"testing"

	"github.com/bcho/annie/pkg/jsonconfig"
)

func TestBasic(t *testing.T) {
	config, _ := jsonconfig.LoadFromString("{}")

	SetupMemoryManagerWithConfig(config)
	manager := GetMemoryManager()

	if manager == nil {
		t.Errorf("cannot create manager")
	}

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
