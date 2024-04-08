package persistence

import (
	"encoding/binary"

	badger "github.com/dgraph-io/badger/v4"
)

const (
	SizeOfLogFile = 64 * 1024 * 1024
)

type BadgerRepository struct {
	db *badger.DB
}

func New(contextPath string) *BadgerRepository {
	r := BadgerRepository{}

	var err error
	options := badger.
		DefaultOptions(contextPath).
		WithLogger(nil).
		WithValueLogFileSize(SizeOfLogFile)
	r.db, err = badger.Open(options)
	if err != nil {
		panic(err)
	}

	return &r
}

func (r *BadgerRepository) ReadUInt64(key string) (uint64, error) {
	var value uint64

	err := r.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err == nil {
			err = item.Value(func(val []byte) error {
				value = binary.LittleEndian.Uint64(val)
				return nil
			})
		}

		return err
	})

	return value, err
}

func (r *BadgerRepository) WriteUint64(key string, value uint64) error {
	err := r.db.Update(func(txn *badger.Txn) error {
		bytes := binary.LittleEndian.AppendUint64([]byte{}, value)
		entry := badger.NewEntry([]byte(key), bytes)
		err := txn.SetEntry(entry)
		return err
	})

	return err
}
