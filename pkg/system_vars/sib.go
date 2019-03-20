package sv

// SIBScriptName is the name of the SIBScript system variable
//
// This sytem variable stores the chaincode script for calculating SIB.
const SIBScriptName = "SIBScript"

// SIBScriptDefault exists in case the system variable isn't present on the
// blockchain.
// See https://github.com/oneiro-ndev/chaincode_scripts/blob/e8289c66fd39b0830cbc06066f771d8eafead370/src/sib/sib.chasm
const SIBScriptDefault = "oAAmABCl1OgADwJGBSYAnGkw3QDDiiAQjwUlAIhSanTBiiUAiFJqdBCPJQCIUmp0CSUAiFJqdEElAIhSanRJJgCcaTDdACUAiFJqdEFGQIg="

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
