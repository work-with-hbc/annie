/*
Testcases for jsonconfig.
*/

package jsonconfig

import (
	"testing"
)

func TestNormalConfig(t *testing.T) {
	config, err := LoadFromFile("testdata/normal.json")
	if err != nil {
		t.Fatal(err)
	}

	versionConfig := config.GetConfigObj("version")
	if versionConfig == nil {
		t.Errorf("versionConfig is nil")
	}

	if majorVersion := versionConfig.GetInt("major"); majorVersion != 0 {
		t.Errorf("majorVersion: expected 0, got %d", majorVersion)
	}

	if minorVersion := versionConfig.GetFloat("minor"); minorVersion != 0.1 {
		t.Errorf("minorVersion: expected 0.1, got %f", minorVersion)
	}

	if stable := config.GetBool("stable"); stable != false {
		t.Errorf("stable: expecected false, got true")
	}
}

func TestBrokenConfig(t *testing.T) {
	_, err := LoadFromFile("testdata/broken.json")
	if err == nil {
		t.Errorf("should not parse broken json")
	}
}

func TestGetConfig(t *testing.T) {
	config, err := LoadFromFile("testdata/nested.json")
	if err != nil {
		t.Fatal(err)
	}

	if intVal := config.GetInt("int.nested"); intVal != 42 {
		t.Errorf("nested int: expected 42, got %d", intVal)
	}

	if floatVal := config.GetFloat("float.nested"); floatVal != 3.14 {
		t.Errorf("nested float: expected 3.14, got %f", floatVal)
	}

	if boolVal := config.GetBool("bool.nested"); boolVal != false {
		t.Errorf("nested bool: expected false, got true")
	}

	if stringVal := config.GetString("string.nested"); stringVal != "jsonconfig" {
		t.Errorf("nested string: expected jsonconfig, got %s", stringVal)
	}

	if subConfig := config.GetConfigObj("config.nested"); subConfig == nil {
		t.Errorf("nested subconfig: got nil")
	}
}

func TestGetArrayConfigObj(t *testing.T) {
	configs, err := LoadFromFile("testdata/array_configs.json")
	if err != nil {
		t.Fatal(err)
	}

	subConfigs := configs.GetArrayConfigObj("sub_configs")
	if subConfigs == nil {
		t.Errorf("cannot get array config obj")
	}

	if len(subConfigs) != configs.GetInt("sub_configs_count") {
		t.Errorf("get array config obj failed: %q", subConfigs)
	}
}

func TestGetDefaultValue(t *testing.T) {
	configs, err := LoadFromString("{}")
	if err != nil {
		t.Fatal(err)
	}

	if x := configs.GetDefaultString("not_existed_key", "42"); x != "42" {
		t.Errorf("get default string failed, expected 42, got %s", x)
	}

	if x := configs.GetDefaultInt("not_existed_key", 42); x != 42 {
		t.Errorf("get default int failed, expected 42, got %d", x)
	}

	if x := configs.GetDefaultFloat("not_existed_key", 42.0); x != 42.0 {
		t.Errorf("get default int failed, expected 42.0, got %f", x)
	}

	if x := configs.GetDefaultBool("not_existed_key", true); !x {
		t.Errorf("get default int failed, expected true, got false")
	}
}
