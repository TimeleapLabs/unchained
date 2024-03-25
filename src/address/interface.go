package address

import (
	"bytes"
	"encoding/hex"

	"github.com/KenshiTech/unchained/xerrors"
	"golang.org/x/crypto/sha3"
)

type Address interface {
	// Hex returns the hex representation of the address with the hex prefix.
	// old version: x,_ :=  CalculateHex(input []byte) (string, [20]byte)
	Hex() string

	// Strings returns the internal base32 formatted of the address including checksum.
	// old version: equivalents to Calculate()
	String() string

	// Raw returns the raw representation of the address [20]bytes.
	// old version: _,x :=  CalculateHex(input []byte) (string, [20]byte)
	Raw() [addressLength]byte
}

type address struct {

	// pk is the raw public key
	// TODO more specification of how address is encoded required, is it hex or raw bytes or etc.
	pk                []byte
	calculatedAddress []byte
	checksum          string
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

var _ Address = &address{}

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
