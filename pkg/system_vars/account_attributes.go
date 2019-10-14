package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

import (
	"github.com/oneiro-ndev/ndaumath/pkg/address"
)

//go:generate msgp -io=0

// AccountAttributesName is the name of the AccountAttributes system variable
//
// This sytem variable stores a map[string]map[string]struct{} which maps account addresses to
// a map of attributes associated with that account. For example, if AccountAttributes[addr]["x"]
// exists, then the account with address 'addr' is an exchange account.
const AccountAttributesName = "AccountAttributes"

// Available account attributes.
const (
	AccountAttributeExchange string = "x"
)

// AccountAttributes is a list of EAI fees and their destinations
type AccountAttributes map[string]map[string]struct{}

// Zeroize implements validatable
func (aa *AccountAttributes) Zeroize() {
	*aa = make(AccountAttributes)
}

// Validate implements SelfValidatable
//
// In this case, it exists to reduce the chance that someone reversed the order
// of the map keys by enforcing that all top-level keys must be valid addresses
func (aa *AccountAttributes) Validate() bool {
	for addr := range *aa {
		_, err := address.Validate(addr)
		if err != nil {
			return false
		}
	}
	return true
}

func init() {
	RegisterTypeValidator(AccountAttributesName, &AccountAttributes{})
}
