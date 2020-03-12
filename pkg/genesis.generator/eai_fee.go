package generator

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


import (
	"github.com/ndau/ndaumath/pkg/address"
	"github.com/ndau/ndaumath/pkg/constants"
	"github.com/ndau/ndaumath/pkg/signature"
	"github.com/ndau/ndaumath/pkg/signed"
	math "github.com/ndau/ndaumath/pkg/types"
	sv "github.com/ndau/system_vars/pkg/system_vars"
	"github.com/pkg/errors"
)

func makeEAIFeeTable() (table sv.EAIFeeTable, err error) {
	add := func(thousandths int64, maker func(int64) (sv.EAIFee, error)) {
		var fee sv.EAIFee
		fee, err = maker(thousandths)
		table = append(table, fee)
	}

	// ndev operations
	add(40, makeEAIFee)
	if err != nil {
		return
	}

	// ntrd operations
	add(10, makeEAIFee)
	if err != nil {
		return
	}

	// rfe acct
	add(1, makeEAIFee)
	if err != nil {
		return
	}

	// rewards nomination acct
	add(1, makeEAIFee)
	if err != nil {
		return
	}

	// node rewards
	add(98, makeNodeRewardEAIFee)
	if err != nil {
		return
	}

	return
}

func makeEAIFee(thousandths int64) (ef sv.EAIFee, err error) {
	var public signature.PublicKey
	public, _, err = signature.Generate(signature.Ed25519, nil)
	if err != nil {
		return
	}
	var addr address.Address
	addr, err = address.Generate(address.KindNdau, public.KeyBytes())
	ef.To = &addr
	if err != nil {
		return
	}

	var fee int64
	fee, err = signed.MulDiv(thousandths, constants.QuantaPerUnit, 1000)
	ef.Fee = math.Ndau(fee)
	if err != nil {
		return
	}

	return
}

func makeNodeRewardEAIFee(thousandths int64) (sv.EAIFee, error) {
	fee, err := signed.MulDiv(thousandths, constants.QuantaPerUnit, 1000)
	return sv.EAIFee{
		Fee: math.Ndau(fee),
		To:  nil,
	}, errors.Wrap(err, "making node rewards eai fee")
}
