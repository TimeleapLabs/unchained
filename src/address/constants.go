package address

import "encoding/base32"

const (
	hexPrefix = "0x"

	base32Chars = "0123456789ABCDEFGHJKMNPQRSTUVXYZ"

	addressLength = 20
)

var InternalBase32Encoding = base32.NewEncoding(base32Chars)
