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
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/stretchr/testify/require"
)

func TestValueRoundtrip(t *testing.T) {
	addr, err := address.Generate(address.KindUser, []byte("a test address"))
	require.NoError(t, err)

	const qtyAddrs = 5
	addrs := make([]address.Address, 0, qtyAddrs)
	paddrs := make([]*address.Address, 0, qtyAddrs)
	for i := 0; i < qtyAddrs; i++ {
		addr, err := address.Generate(address.KindUser, []byte(fmt.Sprintf("another test address: %d", i)))
		require.NoError(t, err)
		addrs = append(addrs, addr)
		paddrs = append(paddrs, &addr)
	}

	type testcase struct {
		name    string
		nonzero interface{}
	}

	cases := []testcase{
		{"bool", true},
		{"int", 1},
		{"int64", int64(2)},
		{"uint", uint(3)},
		{"uint64", uint64(4)},
		{"string", "5"},
		{"time.Time", time.Now()},
		{"[]uint8", []byte{0x67, 0x89, 0xab}}, // synonym for []byte
		{"*address.Address", &addr},
		{"address.Address", addr},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := PackValue(tc.nonzero)
			require.NoError(t, err)

			require.Equal(t, unpointerTypename(tc.name), value.GoType)

			unpacked, err := value.Unpack()
			require.NoError(t, err)
			if _, ok := unpacked.(Valuable); ok {
				require.Equal(t, pointerize(tc.nonzero), unpacked)
			} else {
				require.Equal(t, tc.nonzero, unpacked)
			}
		})
	}
	// list of address requires special handling in the final test
	cases = []testcase{
		{"[]*address.Address", paddrs},
		{"[]address.Address", addrs},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := PackValue(tc.nonzero)
			require.NoError(t, err)

			require.Equal(t, unpointerTypename(tc.name), value.GoType)

			unpacked, err := value.Unpack()
			require.NoError(t, err)

			ua, ok := unpacked.([]Valuable)
			require.True(t, ok)

			nzv := reflect.ValueOf(tc.nonzero)
			require.Equal(t, nzv.Len(), len(ua))
			for idx := 0; idx < len(ua); idx++ {
				addr, ok := ua[idx].(*address.Address)
				require.True(t, ok)

				expect, ok := pointerize(nzv.Index(idx).Interface()).(*address.Address)
				require.True(t, ok, "expect must have addresses")

				require.Equal(t, expect.String(), addr.String())
			}
		})
	}
}
