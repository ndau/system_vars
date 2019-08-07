package sv

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
