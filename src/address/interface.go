package address

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
