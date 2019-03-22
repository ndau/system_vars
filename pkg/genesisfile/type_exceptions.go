package genesisfile

// sometimes we define a convenience type in a library which breaks
// our code. The _right_ answer to this problem would be to recursively
// unwrap typedefs, trying each of them until we found one we could handle,
// a primitive, or a structdef
//
// we're not doing that
//
// the _expedient_ answer is to create a map of convenience typenames, and
// what we unpack them to

var typeExceptions map[string]string

func init() {
	typeExceptions = map[string]string{
		"eai.RateTable":  "[]eai.RTRow",
		"sv.EAIFeeTable": "[]sv.EAIFee",
	}
}
