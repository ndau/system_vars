package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

// ChangeSchemaAddressName is the name of the ChangeSchemaAddress system variable
//
// The value contained in this system variable must be of type address.Address
const ChangeSchemaAddressName = "ChangeSchemaAddress"

// ChangeSchemaOwnershipName is the name of the public ownership key
const ChangeSchemaOwnershipName = "ChangeSchemaOwnership"

// ChangeSchemaOwnershipPrivateName is the name of the private ownership key
const ChangeSchemaOwnershipPrivateName = "ChangeSchemaOwnershipPrivate"

// ChangeSchemaValidationName is the name of the public validation key
const ChangeSchemaValidationName = "ChangeSchemaValidation"

// ChangeSchemaValidationPrivateName is the name of the private validation key
const ChangeSchemaValidationPrivateName = "ChangeSchemaValidationPrivate"

// ChangeSchema encapsulates data about the ChangeSchema system variables in a structured way.
var ChangeSchema = SysAcct{
	Name:    "ChangeSchema",
	Address: ChangeSchemaAddressName,
	Ownership: Keypair{
		Public:  ChangeSchemaOwnershipName,
		Private: ChangeSchemaOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  ChangeSchemaValidationName,
		Private: ChangeSchemaValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(ChangeSchemaAddressName, ValidateAddress)
}
