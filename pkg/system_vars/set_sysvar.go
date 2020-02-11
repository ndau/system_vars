package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// SetSysvarAddressName is the name of the SetSysvarAddress system variable
//
// The value contained in this system variable must be of type address.Address
const SetSysvarAddressName = "SetSysvarAddress"

// SetSysvarOwnershipName is the name of the public ownership key
const SetSysvarOwnershipName = "SetSysvarOwnership"

// SetSysvarOwnershipPrivateName is the name of the private ownership key
const SetSysvarOwnershipPrivateName = "SetSysvarOwnershipPrivate"

// SetSysvarValidationName is the name of the public validation key
const SetSysvarValidationName = "SetSysvarValidation"

// SetSysvarValidationPrivateName is the name of the private validation key
const SetSysvarValidationPrivateName = "SetSysvarValidationPrivate"

// SetSysvar encapsulates data about the SetSysvar system variables in a structured way.
var SetSysvar = SysAcct{
	Name:    "SetSysvar",
	Address: SetSysvarAddressName,
	Ownership: Keypair{
		Public:  SetSysvarOwnershipName,
		Private: SetSysvarOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  SetSysvarValidationName,
		Private: SetSysvarValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(SetSysvarAddressName, ValidateAddress)
}
