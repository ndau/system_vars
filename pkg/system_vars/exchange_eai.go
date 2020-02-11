package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// ExchangeEAIScriptName is the name of the ExchangeEAIScript system variable
//
// This sytem variable stores the chaincode script for calculating EAI rate for
// exchange accounts.
const ExchangeEAIScriptName = "ExchangeEAIScript"

func init() {
	RegisterFuncValidator(ExchangeEAIScriptName, ValidateChaincode)
}
