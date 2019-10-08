package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

// SIBScriptName is the name of the SIBScript system variable
//
// This sytem variable stores the chaincode script for calculating SIB.
const SIBScriptName = "SIBScript"

func init() {
	RegisterFuncValidator(SIBScriptName, ValidateChaincode)
}

// RecordPriceAddressName is the name of the RecordPriceAddress system variable
//
// The value contained in this system variable must be of type address.Address
const RecordPriceAddressName = "RecordPriceAddress"

// RecordPriceOwnershipName is the name of the public ownership key
const RecordPriceOwnershipName = "RecordPriceOwnership"

// RecordPriceOwnershipPrivateName is the name of the private ownership key
const RecordPriceOwnershipPrivateName = "RecordPriceOwnershipPrivate"

// RecordPriceValidationName is the name of the public validation key
const RecordPriceValidationName = "RecordPriceValidation"

// RecordPriceValidationPrivateName is the name of the private validation key
const RecordPriceValidationPrivateName = "RecordPriceValidationPrivate"

// RecordPrice encapsulates data about the RecordPrice system variables in a structured way.
var RecordPrice = SysAcct{
	Name:    "RecordPrice",
	Address: RecordPriceAddressName,
	Ownership: Keypair{
		Public:  RecordPriceOwnershipName,
		Private: RecordPriceOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  RecordPriceValidationName,
		Private: RecordPriceValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(RecordPriceAddressName, ValidateAddress)
}
