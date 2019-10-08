package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

// BPCRulesAccountAddressName is the name of the BPCRulesAccountAddress system variable
//
// The value contained in this system variable must be of type address.Address
const BPCRulesAccountAddressName = "BPCRulesAccountAddress"

// BPCRulesAccountOwnershipName is the name of the public ownership key
const BPCRulesAccountOwnershipName = "BPCRulesAccountOwnership"

// BPCRulesAccountOwnershipPrivateName is the name of the private ownership key
const BPCRulesAccountOwnershipPrivateName = "BPCRulesAccountOwnershipPrivate"

// BPCRulesAccountValidationName is the name of the public validation key
const BPCRulesAccountValidationName = "BPCRulesAccountValidation"

// BPCRulesAccountValidationPrivateName is the name of the private validation key
const BPCRulesAccountValidationPrivateName = "BPCRulesAccountValidationPrivate"

// BPCRulesAccount encapsulates data about the BPCRulesAccount system variables in a structured way.
var BPCRulesAccount = SysAcct{
	Name:    "BPCRulesAccount",
	Address: BPCRulesAccountAddressName,
	Ownership: Keypair{
		Public:  BPCRulesAccountOwnershipName,
		Private: BPCRulesAccountOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  BPCRulesAccountValidationName,
		Private: BPCRulesAccountValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(BPCRulesAccountAddressName, ValidateAddress)
}
