package frost

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestFilterOnlineSigners(t *testing.T) {
	type args struct {
		signers []common.Address
		onlines []string
	}
	tests := []struct {
		name string
		args args
		want []common.Address
	}{
		{
			name: "Test case 1",
			args: args{
				signers: []common.Address{
					common.HexToAddress("0x1"),
					common.HexToAddress("0x2"),
					common.HexToAddress("0x3"),
				},
				onlines: []string{
					"0x0000000000000000000000000000000000000001",
					"0x0000000000000000000000000000000000000003",
				},
			},
			want: []common.Address{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x3"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				signers: []common.Address{
					common.HexToAddress("0x1"),
					common.HexToAddress("0x2"),
					common.HexToAddress("0x3"),
				},
				onlines: []string{
					"0x0000000000000000000000000000000000000001",
					"0x0000000000000000000000000000000000000002",
					"0x0000000000000000000000000000000000000003",
				},
			},
			want: []common.Address{
				common.HexToAddress("0x1"),
				common.HexToAddress("0x2"),
				common.HexToAddress("0x3"),
			},
		},
		{
			name: "Test case 3",
			args: args{
				signers: []common.Address{
					common.HexToAddress("0x1"),
					common.HexToAddress("0x2"),
					common.HexToAddress("0x3"),
				},
				onlines: []string{
					"0x0000000000000000000000000000000000000002",
					"0x0000000000000000000000000000000000000003",
				},
			},
			want: []common.Address{
				common.HexToAddress("0x2"),
				common.HexToAddress("0x3"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterOnlineSigners(tt.args.signers, tt.args.onlines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterOnlineSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
