package sv

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
