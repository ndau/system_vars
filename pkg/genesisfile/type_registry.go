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
	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/oneiro-ndev/ndaumath/pkg/eai"
	"github.com/oneiro-ndev/ndaumath/pkg/signature"
	sv "github.com/oneiro-ndev/system_vars/pkg/system_vars"
)

var valuableRegistry map[string]Valuable

func init() {
	valuableRegistry = make(map[string]Valuable)

	RegisterValuable(&address.Address{})
	RegisterValuable(&signature.PublicKey{})
	RegisterValuable(&eai.RTRow{})
	RegisterValuable(&sv.EAIFee{})
}

// RegisterValuable registers a Valuable instance
//
// Valuables can only be unmarshalled if they have been registered. It is a good
// idea to put a bunch of RegisterValuable calls in in an init function in your
// code in order to be able to unmarshal your types.
//
// Pre-registered types:
// - address.Address
// - signature.PublicKey
// - eai.RTRow
func RegisterValuable(v Valuable) {
	valuableRegistry[getTypeName(v)] = emptyCopy(v)
}
