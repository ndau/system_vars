package system_vars

import (
	"encoding"
	"encoding/base64"

	"github.com/oneiro-ndev/ndaumath/pkg/address"
	math "github.com/oneiro-ndev/ndaumath/pkg/types"
	"github.com/pkg/errors"
)

//go:generate msgp -io=0

// EAIFeeTableName names the EAI fee table
//
// The system variable of this name must have the type EAIFeeTable
const EAIFeeTableName = "EAIFeeTable"

// EAIFeeTable is a list of EAI fees and their destinations
type EAIFeeTable []EAIFee

// An EAIFee is a fee applied to accrued EAI when crediting.
//
// The fee is listed as Ndau; the listed value is multiplied by the number
// of Ndau actually earned as EAI.
//
// The fee is credited to the account at the listed address. If the destination
// is nil, it is considered to be a node reward, and is tracked in internal
// state instead of going into an account.
type EAIFee struct {
	Fee math.Ndau
	To  *address.Address
}

var _ encoding.TextMarshaler = (*EAIFee)(nil)
var _ encoding.TextUnmarshaler = (*EAIFee)(nil)

// MarshalText implements encoding.TextMarshaler
func (f EAIFee) MarshalText() ([]byte, error) {
	bytes, err := f.MarshalMsg(nil)
	if err != nil {
		return bytes, errors.Wrap(err, "marshalling bytes")
	}
	return []byte(base64.StdEncoding.EncodeToString(bytes)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (f *EAIFee) UnmarshalText(text []byte) error {
	bytes, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return errors.Wrap(err, "decoding b64")
	}
	_, err = f.UnmarshalMsg(bytes)
	if err != nil {
		return errors.Wrap(err, "unmarshalling bytes")
	}

	return nil
}
