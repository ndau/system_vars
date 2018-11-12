package sv

// MinDurationBetweenNodeRewardNominationsName names the minimum duration
// permitted between node rewards nominations
//
// The system variable named by this must have the type math.Duration
const MinDurationBetweenNodeRewardNominationsName = "MinDurationBetweenNodeRewardNominations"

// NominateNodeRewardAddressName is the name of the NominateNodeRewardAddress system variable
//
// The value contained in this system variable must be of type address.Address
const NominateNodeRewardAddressName = "NominateNodeRewardAddress"

// NominateNodeRewardOwnershipName is the name of the public ownership key
const NominateNodeRewardOwnershipName = "NominateNodeRewardOwnership"

// NominateNodeRewardOwnershipPrivateName is the name of the private ownership key
const NominateNodeRewardOwnershipPrivateName = "NominateNodeRewardOwnershipPrivate"

// NominateNodeRewardValidationName is the name of the public validation key
const NominateNodeRewardValidationName = "NominateNodeRewardValidation"

// NominateNodeRewardValidationPrivateName is the name of the private validation key
const NominateNodeRewardValidationPrivateName = "NominateNodeRewardValidationPrivate"

// NodeRewardNominationTimeoutName names the maximum interval permitted between
// valid NominateNodeReward and ClaimNodeReward transactions.
//
// The system variable named by this must have the type math.Duration
const NodeRewardNominationTimeoutName = "NodeRewardNominationTimeout"

// NominateNodeReward encapsulates data about the NominateNodeReward system variables in a structured way.
var NominateNodeReward = SysAcct{
	Name:    "NominateNodeReward",
	Address: NominateNodeRewardAddressName,
	Ownership: Keypair{
		Public:  NominateNodeRewardOwnershipName,
		Private: NominateNodeRewardOwnershipPrivateName,
	},
	Validation: Keypair{
		Public:  NominateNodeRewardValidationName,
		Private: NominateNodeRewardValidationPrivateName,
	},
}
