package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

// DisputeRulesAccountAddressName is the name of the DisputeRulesAccountAddress system variable
//
// The value contained in this system variable must be of type address.Address
const DisputeRulesAccountAddressName = "DisputeRulesAccountAddress"

// DisputeRulesAccountOwnershipName is the name of the public ownership key
const DisputeRulesAccountOwnershipName = "DisputeRulesAccountOwnership"

// DisputeRulesAccountOwnershipPrivateName is the name of the private ownership key
const DisputeRulesAccountOwnershipPrivateName = "DisputeRulesAccountOwnershipPrivate"

// DisputeRulesAccountValidationName is the name of the public validation key
const DisputeRulesAccountValidationName = "DisputeRulesAccountValidation"

// DisputeRulesAccountValidationPrivateName is the name of the private validation key
const DisputeRulesAccountValidationPrivateName = "DisputeRulesAccountValidationPrivate"

// DisputeRulesAccount encapsulates data about the DisputeRulesAccount system variables in a structured way.
var DisputeRulesAccount = SysAcct{
	Name:    "DisputeRulesAccount",
	Address: DisputeRulesAccountAddressName,
	Ownership: Keypair{
		Public:  DisputeRulesAccountOwnershipName,
		Private: DisputeRulesAccountOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  DisputeRulesAccountValidationName,
		Private: DisputeRulesAccountValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(DisputeRulesAccountAddressName, ValidateAddress)
}
