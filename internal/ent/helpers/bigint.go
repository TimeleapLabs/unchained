package helpers

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

const (
	BaseOfNumbers = 10
)

func (b *BigInt) Scan(src any) error {
	var i sql.NullString
	if err := i.Scan(src); err != nil {
		return err
	}
	if !i.Valid {
		return nil
	}
	if _, ok := b.Int.SetString(i.String, BaseOfNumbers); ok {
		return nil
	}
	return fmt.Errorf("could not scan type %T with value %v into BigInt", src, src)
}

func (b *BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}
