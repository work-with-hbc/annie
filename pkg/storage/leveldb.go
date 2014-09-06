/*
Leveldb storage support.
*/

package storage

import (
	"log"

	"github.com/bcho/annie/pkg/jsonconfig"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBStorage struct {
	db *leveldb.DB
}

func (s LevelDBStorage) Save(key string, thing []byte) error {
	return s.db.Put([]byte(key), thing, nil)
}

func (s LevelDBStorage) Get(key string) ([]byte, error) {
	data, err := s.db.Get([]byte(key), nil)

	return data, err
}

func (s LevelDBStorage) Has(key string) bool {
	_, err := s.db.Get([]byte(key), nil)

	return err == nil
}

func (s LevelDBStorage) Forget(key string) error {
	return s.db.Delete([]byte(key), nil)
}

func (s *LevelDBStorage) Terminate() error {
	s.db.Close()

	return nil
}

func makeLevelDBStorage(dbPath string) *LevelDBStorage {
	s := new(LevelDBStorage)
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}
	s.db = db

	return s
}

func MakeLevelDBStorageWithConfig(config *jsonconfig.Config) Storage {
	dbPath := config.GetString("path")

	log.Printf("Initializing leveldb: %s", dbPath)

	return makeLevelDBStorage(dbPath)
}
