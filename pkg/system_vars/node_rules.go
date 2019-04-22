package sv

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
