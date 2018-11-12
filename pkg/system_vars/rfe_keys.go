package sv

// ReleaseFromEndowmentAddressName is the name of the ReleaseFromEndowmentAddress system variable
//
// The value contained in this system variable must be of type address.Address
const ReleaseFromEndowmentAddressName = "ReleaseFromEndowmentAddress"

// ReleaseFromEndowmentOwnershipName is the name of the public ownership key
const ReleaseFromEndowmentOwnershipName = "ReleaseFromEndowmentOwnership"

// ReleaseFromEndowmentOwnershipPrivateName is the name of the private ownership key
const ReleaseFromEndowmentOwnershipPrivateName = "ReleaseFromEndowmentOwnershipPrivate"

// ReleaseFromEndowmentValidationName is the name of the public validation key
const ReleaseFromEndowmentValidationName = "ReleaseFromEndowmentValidation"

// ReleaseFromEndowmentValidationPrivateName is the name of the private validation key
const ReleaseFromEndowmentValidationPrivateName = "ReleaseFromEndowmentValidationPrivate"

// ReleaseFromEndowment encapsulates data about the ReleaseFromEndowment system variables in a structured way.
var ReleaseFromEndowment = SysAcct{
	Name: "ReleaseFromEndowment",
	Ownership: Keypair{
		Public:  ReleaseFromEndowmentOwnershipName,
		Private: ReleaseFromEndowmentOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  ReleaseFromEndowmentValidationName,
		Private: ReleaseFromEndowmentValidationPrivateName,
	},
}
