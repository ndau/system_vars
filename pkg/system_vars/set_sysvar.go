package sv

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
