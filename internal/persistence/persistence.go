package persistence

import (
	"encoding/binary"

	badger "github.com/dgraph-io/badger/v4"
)

var DB *badger.DB

const (
	SizeOfLogFile = 64 * 1024 * 1024
)

func Start(contextPath string) {
	var err error
	options := badger.
		DefaultOptions(contextPath).
		WithLogger(nil).
		WithValueLogFileSize(SizeOfLogFile)
	DB, err = badger.Open(options)
	if err != nil {
		panic(err)
	}
}

func ReadUInt64(key string) (uint64, error) {
	var value uint64

	err := DB.View(func(txn *badger.Txn) error {
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

func WriteUint64(key string, value uint64) error {
	err := DB.Update(func(txn *badger.Txn) error {
		bytes := binary.LittleEndian.AppendUint64([]byte{}, value)
		entry := badger.NewEntry([]byte(key), bytes)
		err := txn.SetEntry(entry)
		return err
	})

	return err
}
