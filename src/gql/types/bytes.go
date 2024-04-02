package types

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/KenshiTech/unchained/src/log"
)

type Bytes []byte

// UnmarshalGQL implements the graphql.Unmarshaler interface.
func (bytes *Bytes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("Bytes must be a string")
	}
	decoded, err := hex.DecodeString(str)
	if err != nil {
		return fmt.Errorf("Bytes must be a hex string")
	}
	*bytes = decoded
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface.
func (bytes Bytes) MarshalGQL(w io.Writer) {
	hexValue := fmt.Sprintf(`"%x"`, bytes)
	_, err := w.Write([]byte(hexValue))
	if err != nil {
		log.Logger.Error(err.Error())
	}
}
