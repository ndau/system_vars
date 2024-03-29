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
	"sort"

	"github.com/ndau/chaincode/pkg/vm"
	"github.com/ndau/msgp-well-known-types/wkt"
	"github.com/ndau/ndaumath/pkg/address"
	math "github.com/ndau/ndaumath/pkg/types"
	"github.com/tinylib/msgp/msgp"
)

var funcValidators map[string]func([]byte) bool
var typeValidators map[string]Validatable

// RegisterFuncValidator registers a function validator for a sysvar
func RegisterFuncValidator(name string, v func([]byte) bool) {
	if funcValidators == nil {
		funcValidators = make(map[string]func([]byte) bool)
	}
	funcValidators[name] = v
}

// FuncValidators returns the list of known function validators
func FuncValidators() []string {
	fvs := make([]string, 0, len(funcValidators))
	for fv := range funcValidators {
		fvs = append(fvs, fv)
	}
	sort.Strings(fvs)
	return fvs
}

// RegisterTypeValidator registers a type validator for a sysvar
func RegisterTypeValidator(name string, v Validatable) {
	if typeValidators == nil {
		typeValidators = make(map[string]Validatable)
	}
	typeValidators[name] = v
}

// TypeValidators returns the list of know type validators
func TypeValidators() []string {
	tvs := make([]string, 0, len(typeValidators))
	for tv := range typeValidators {
		tvs = append(tvs, tv)
	}
	sort.Strings(tvs)
	return tvs
}

// HasValidator is true when a validator exists for the given sysvar
func HasValidator(name string) bool {
	_, okf := funcValidators[name]
	_, okt := typeValidators[name]
	return okf || okt
}

// IsValid is true when the supplied data is a valid instance of the given
// system variable.
//
// In the event the name is unknown, this function will return nil.
func IsValid(name string, data []byte) (*bool, error) {
	if fv, ok := funcValidators[name]; ok {
		v := fv(data)
		return &v, nil
	}
	if tv, ok := typeValidators[name]; ok {
		tv.Zeroize()
		leftovers, err := tv.UnmarshalMsg(data)
		v := err == nil && len(leftovers) == 0
		if v {
			if sv, ok := tv.(SelfValidatable); ok {
				v = sv.Validate()
			}
		}
		return &v, err
	}
	return nil, nil
}

// A Validatable type is one which can be unmarshaled and zeroized
type Validatable interface {
	msgp.Unmarshaler

	// reset this instance to the zero value
	Zeroize()
}

// A SelfValidatable type is one which applies additional self-validation
type SelfValidatable interface {
	Validate() bool
}

func validateM(m msgp.Unmarshaler, data []byte) bool {
	l, err := m.UnmarshalMsg(data)
	return err == nil && len(l) == 0
}

// ValidateDuration ensures this value works as a Duration
func ValidateDuration(data []byte) bool {
	d := math.Duration(0)
	return validateM(&d, data)
}

// ValidateNdau ensures this value works as Ndau
func ValidateNdau(data []byte) bool {
	n := math.Ndau(0)
	return validateM(&n, data)
}

// ValidateAddress ensures this value works as an Address
func ValidateAddress(data []byte) bool {
	a := address.Address{}
	v := validateM(&a, data)
	if v {
		v = a.Revalidate() == nil
	}
	return v
}

// ValidateChaincode ensures this value works as a Chaincode script
func ValidateChaincode(data []byte) bool {
	b := wkt.Bytes{}
	v := validateM(&b, data)
	if v {
		c := vm.ToChaincode(b)
		v = c.IsValid() == nil
	}
	return v
}

// ValidateUInt64 ensures this value works as a wkt.Uint64
func ValidateUInt64(data []byte) bool {
	i := wkt.Uint64(0)
	return validateM(&i, data)
}
