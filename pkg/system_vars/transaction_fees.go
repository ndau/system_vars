package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// TxFeeScriptName is the name of the TxFeeScript system variable
//
// This sytem variable store the chaincode script used to calculate transaction
// fees.
const TxFeeScriptName = "TransactionFeeScript"

func init() {
	RegisterFuncValidator(TxFeeScriptName, ValidateChaincode)
}
