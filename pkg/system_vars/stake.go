package sv

// MinNodeRegistrationStakeName names the MinNodeRegistrationStake system variable
const MinNodeRegistrationStakeName = "MinNodeRegistrationStakeAmount"

func init() {
	RegisterFuncValidator(MinNodeRegistrationStakeName, ValidateNdau)
}
