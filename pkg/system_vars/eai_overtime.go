package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

// EAIOvertime names the EAIOvertime system variable
//
// EAI Overtime is a math.Duration constraining the max awarded interval between
// CreditEAI transactions.
const EAIOvertime = "EAIOvertime"

func init() {
	RegisterFuncValidator(EAIOvertime, ValidateDuration)
}
