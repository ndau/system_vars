package sv

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


// NodeRulesAccountAddressName is the name of the NodeRulesAccountAddress system variable
//
// The value contained in this system variable must be of type address.Address
const NodeRulesAccountAddressName = "NodeRulesAccountAddress"

// NodeRulesAccountOwnershipName is the name of the public ownership key
const NodeRulesAccountOwnershipName = "NodeRulesAccountOwnership"

// NodeRulesAccountOwnershipPrivateName is the name of the private ownership key
const NodeRulesAccountOwnershipPrivateName = "NodeRulesAccountOwnershipPrivate"

// NodeRulesAccountValidationName is the name of the public validation key
const NodeRulesAccountValidationName = "NodeRulesAccountValidation"

// NodeRulesAccountValidationPrivateName is the name of the private validation key
const NodeRulesAccountValidationPrivateName = "NodeRulesAccountValidationPrivate"

// NodeRulesAccount encapsulates data about the NodeRulesAccount system variables in a structured way.
var NodeRulesAccount = SysAcct{
	Name:    "NodeRulesAccount",
	Address: NodeRulesAccountAddressName,
	Ownership: Keypair{
		Public:  NodeRulesAccountOwnershipName,
		Private: NodeRulesAccountOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  NodeRulesAccountValidationName,
		Private: NodeRulesAccountValidationPrivateName,
	},
}

func init() {
	RegisterFuncValidator(NodeRulesAccountAddressName, ValidateAddress)
}
