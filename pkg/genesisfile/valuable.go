package genesisfile

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
