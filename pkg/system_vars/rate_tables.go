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
	"github.com/oneiro-ndev/ndaumath/pkg/eai"
)

// UnlockedRateTableName is the name of the UnlockedRateTable system variable.
//
// The system variable of this name has the type eai.RateTable.
//
// This variable defines the basic table of unlocked EAI rates.
const UnlockedRateTableName = "UnlockedRateTable"

// LockedRateTableName is the name of the LockedRateTable system variable.
//
// The system variable of this name has the type eai.RateTable.
//
// This variable defines the basic table of unlocked EAI rates.
const LockedRateTableName = "LockedRateTable"

// ValidateRateTable validates an eai.RateTable
func ValidateRateTable(data []byte) bool {
	rt := eai.RateTable{}
	return validateM(&rt, data)
}

func init() {
	RegisterFuncValidator(UnlockedRateTableName, ValidateRateTable)
	RegisterFuncValidator(LockedRateTableName, ValidateRateTable)
}
