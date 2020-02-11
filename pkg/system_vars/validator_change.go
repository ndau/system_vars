package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// CommandValidatorChangeAddressName is the name of the CommandValidatorChangeAddress system variable
//
// The value contained in this system variable must be of type address.Address
const CommandValidatorChangeAddressName = "CommandValidatorChangeAddress"

// CommandValidatorChangeOwnershipName is the name of the public ownership key
const CommandValidatorChangeOwnershipName = "CommandValidatorChangeOwnership"

// CommandValidatorChangeOwnershipPrivateName is the name of the private ownership key
const CommandValidatorChangeOwnershipPrivateName = "CommandValidatorChangeOwnershipPrivate"

// CommandValidatorChangeValidationName is the name of the public validation key
const CommandValidatorChangeValidationName = "CommandValidatorChangeValidation"

// CommandValidatorChangeValidationPrivateName is the name of the private validation key
const CommandValidatorChangeValidationPrivateName = "CommandValidatorChangeValidationPrivate"

// CommandValidatorChange encapsulates data about the CommandValidatorChange system variables in a structured way.
var CommandValidatorChange = SysAcct{
	Name:    "CommandValidatorChange",
	Address: CommandValidatorChangeAddressName,
	Ownership: Keypair{
		Public:  CommandValidatorChangeOwnershipName,
		Private: CommandValidatorChangeOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  CommandValidatorChangeValidationName,
		Private: CommandValidatorChangeValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(CommandValidatorChangeAddressName, ValidateAddress)
}
