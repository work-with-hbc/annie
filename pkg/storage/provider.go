/*
Storage provider.
*/

package storage

import (
	"github.com/bcho/annie/pkg/jsonconfig"
)

type Storage interface {
	Save(key string, thing []byte) error
	Get(key string) ([]byte, error)
	Has(key string) bool
	Forget(key string) error
	Terminate() error
}

type StorageMaker func(config *jsonconfig.Config) Storage
