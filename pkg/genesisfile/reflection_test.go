package genesisfile

import (
	"reflect"
	"testing"

	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/stretchr/testify/require"
)

func Test_pointerize(t *testing.T) {
	addr, err := address.Generate(address.KindUser, []byte("address for test_pointerize"))
	require.NoError(t, err)

	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want Valuable
	}{
		{"address.Address", args{addr}, Valuable(&addr)},
		{"*address.Address", args{&addr}, Valuable(&addr)},
		{"not valuable", args{"what, really?"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointerize(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointerize() = %v, want %v", got, tt.want)
			}
		})
	}
}
