package address

import (
	"bytes"
	"encoding/hex"

	"github.com/KenshiTech/unchained/xerrors"
	"golang.org/x/crypto/sha3"
)

// address is implementation of the Address interface.
type address struct {
	// pk is the raw public key
	// TODO more specification of how address is encoded required, is it hex or raw bytes or etc.
	pk []byte

	// calculatedAddress is the fixed 20 length raw bytes of the address
	calculatedAddress []byte

	// checksum is string(base32Chars[hashedAddr[0]%32]) + string(base32Chars[hashedAddr[1]%32])
	checksum string
}

var _ Address = &address{}

// NewAddress returns a new Address filled with the given public key.
func NewAddress(bs []byte) (Address, error) {
	// validation
	if len(bs) < 1 {
		return nil, xerrors.ErrNilArgs("address bytes")
	}
	// calculation
	a := new(address)
	a.pk = bytes.Clone(bs)
	{
		//calculate address
		h := sha3.NewShake256()
		h.Write(bs)
		a.calculatedAddress = h.Sum(nil)[:addressLength]
	}
	{
		// calculate checksum
		h := sha3.NewShake256()
		h.Write([]byte(InternalBase32Encoding.EncodeToString(a.calculatedAddress[:addressLength])))
		hbs := h.Sum(nil)
		//todo fix the checksum from protocol or remove this comment,
		// the Base32(hbs[0]%32) != base32Chars[hbs[0]%32]
		// also the checksum is converted byte by byte instead of array of bytes.
		a.checksum = string(base32Chars[hbs[0]%32]) + string(base32Chars[hbs[1]%32])

	}
	return a, nil
}

func (a *address) String() string {
	return InternalBase32Encoding.EncodeToString(a.calculatedAddress) + a.checksum
}

func (a *address) Hex() string {
	return hexPrefix + hex.EncodeToString(a.calculatedAddress[:addressLength])
}

func (a *address) Raw() [addressLength]byte {
	return [addressLength]byte(a.calculatedAddress[:addressLength])
}
