package genesisfile

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


import (
	"reflect"
	"testing"

	"github.com/ndau/ndaumath/pkg/address"
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
