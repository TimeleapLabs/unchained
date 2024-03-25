package address_test

import (
	"fmt"
	"testing"

	"github.com/KenshiTech/unchained/address"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	bs  []byte
	hex string
	raw [20]byte
	str string
}{

	{
		bs:  []byte("hello world"),
		hex: "0x369771bb2cb9d2b04c1d54cca487e372d9f187f7",
		raw: [20]byte{54, 151, 113, 187, 44, 185, 210, 176, 76, 29, 84, 204, 164, 135, 227, 114, 217, 241, 135, 247},
		str: "6TBQ3ESCQ79B0K0XAK6A91Z3EBCZ31ZQF2",
	},
	{
		bs:  []byte("TheUnchainedNetwork"),
		hex: "0x51121847797622891d4b5b1dec9ff9c0186418c1",
		raw: [20]byte{81, 18, 24, 71, 121, 118, 34, 137, 29, 75, 91, 29, 236, 159, 249, 192, 24, 100, 24, 193},
		str: "A491GHUSERH8J7ABBCEYS7ZSR0C686616R",
	},
	// todo add more edge cases
}

func TestAddress(t *testing.T) {
	for _, v := range testCases {
		t.Run(fmt.Sprintf("Case %s", string(v.bs)), func(t *testing.T) {
			addr, err := address.NewAddress(v.bs)
			require.NoError(t, err)
			t.Run("String", func(t *testing.T) {
				require.Equal(t, v.str, addr.String())
			})
			t.Run("Hex", func(t *testing.T) {
				require.Equal(t, v.hex, addr.Hex())
			})
			t.Run("Raw", func(t *testing.T) {
				require.Equal(t, v.raw, addr.Raw())
			})

		})

	}
}
