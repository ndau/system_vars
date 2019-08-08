package sv

// DefaultRecourseDurationName is the name of the DefaultRecourseDuration system variable
const DefaultRecourseDurationName = "DefaultRecourseDuration"

func init() {
	RegisterFuncValidator(DefaultRecourseDurationName, ValidateDuration)
}
