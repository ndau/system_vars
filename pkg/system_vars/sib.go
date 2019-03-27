package sv

// SIBScriptName is the name of the SIBScript system variable
//
// This sytem variable stores the chaincode script for calculating SIB.
const SIBScriptName = "SIBScript"

// RecordPriceAddressName is the name of the RecordPriceAddress system variable
//
// The value contained in this system variable must be of type address.Address
const RecordPriceAddressName = "RecordPriceAddress"

// RecordPriceOwnershipName is the name of the public ownership key
const RecordPriceOwnershipName = "RecordPriceOwnership"

// RecordPriceOwnershipPrivateName is the name of the private ownership key
const RecordPriceOwnershipPrivateName = "RecordPriceOwnershipPrivate"

// RecordPriceValidationName is the name of the public validation key
const RecordPriceValidationName = "RecordPriceValidation"

// RecordPriceValidationPrivateName is the name of the private validation key
const RecordPriceValidationPrivateName = "RecordPriceValidationPrivate"

// RecordPrice encapsulates data about the RecordPrice system variables in a structured way.
var RecordPrice = SysAcct{
	Name:    "RecordPrice",
	Address: RecordPriceAddressName,
	Ownership: Keypair{
		Public:  RecordPriceOwnershipName,
		Private: RecordPriceOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  RecordPriceValidationName,
		Private: RecordPriceValidationPrivateName,
	},
}
