package sv

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

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

const (
	// NodeMaxValidators specifies how many registered nodes are given
	// validation power. The set of nodes is sorted by goodness; only this
	// many of the best receive validation power in proportion to their goodness.
	NodeMaxValidators = "MAX_VALIDATORS"
	// NodeRewardMaxRewarded specifies how many registered nodes are in contention
	// for each node reward. The set of nodes is sorted by goodness; only this
	// many of the best are eligible to receive node rewards.
	NodeRewardMaxRewarded = "MAX_REWARDED"
)

func init() {
	RegisterFuncValidator(NominateNodeRewardAddressName, ValidateAddress)
	RegisterFuncValidator(NodeMaxValidators, ValidateUInt64)
	RegisterFuncValidator(NodeRewardMaxRewarded, ValidateUInt64)
}
