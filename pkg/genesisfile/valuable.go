package genesisfile

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

import (
	"encoding"

	"github.com/tinylib/msgp/msgp"
)

// Valuable is an interface defining what kind of struct can be represented
// in the genesisfile.
//
// TextUnmarshaler is required to read the genesisfile.
// TextMarshaler is required to write the genesisfile.
// msgp.Marshaler is required to write to the chaos chain.
// msgp.Unmarshaler is required to read the value from the chaos chain.
type Valuable interface {
	encoding.TextUnmarshaler
	encoding.TextMarshaler
	msgp.Marshaler
	msgp.Unmarshaler
}
