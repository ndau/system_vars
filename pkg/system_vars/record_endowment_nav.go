package sv

// RecordEndowmentNAVAddressName is the name of the RecordEndowmentNAVAddress system variable
//
// The value contained in this system variable must be of type address.Address
const RecordEndowmentNAVAddressName = "RecordEndowmentNAVAddress"

// RecordEndowmentNAVOwnershipName is the name of the public ownership key
const RecordEndowmentNAVOwnershipName = "RecordEndowmentNAVOwnership"

// RecordEndowmentNAVOwnershipPrivateName is the name of the private ownership key
const RecordEndowmentNAVOwnershipPrivateName = "RecordEndowmentNAVOwnershipPrivate"

// RecordEndowmentNAVValidationName is the name of the public validation key
const RecordEndowmentNAVValidationName = "RecordEndowmentNAVValidation"

// RecordEndowmentNAVValidationPrivateName is the name of the private validation key
const RecordEndowmentNAVValidationPrivateName = "RecordEndowmentNAVValidationPrivate"

// RecordEndowmentNAV encapsulates data about the RecordEndowmentNAV system variables in a structured way.
var RecordEndowmentNAV = SysAcct{
	Name:    "RecordEndowmentNAV",
	Address: RecordEndowmentNAVAddressName,
	Ownership: Keypair{
		Public:  RecordEndowmentNAVOwnershipName,
		Private: RecordEndowmentNAVOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  RecordEndowmentNAVValidationName,
		Private: RecordEndowmentNAVValidationPrivateName,
	},
}
