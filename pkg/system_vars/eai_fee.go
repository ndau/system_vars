package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


import (
	"encoding"
	"fmt"
	"strconv"
	"strings"

	"github.com/ndau/ndaumath/pkg/address"
	math "github.com/ndau/ndaumath/pkg/types"
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
	addr := ""
	if f.To != nil {
		addr = f.To.String()
	}
	return []byte(fmt.Sprintf("%d:%s", f.Fee, addr)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (f *EAIFee) UnmarshalText(text []byte) error {
	parts := strings.Split(string(text), ":")
	if len(parts) != 2 {
		return fmt.Errorf("expected 2 parts; got %d", len(parts))
	}

	nd, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return errors.Wrap(err, "parsing ndau qty as number")
	}
	f.Fee = math.Ndau(nd)

	f.To = nil
	if len(parts[1]) > 0 {
		addr, err := address.Validate(parts[1])
		if err != nil {
			return errors.Wrap(err, "parsing address")
		}
		f.To = &addr
	}

	return nil
}

// Zeroize implements Validatable
func (e *EAIFeeTable) Zeroize() {
	*e = make(EAIFeeTable, 0)
}

func init() {
	RegisterTypeValidator(EAIFeeTableName, &EAIFeeTable{})
}
