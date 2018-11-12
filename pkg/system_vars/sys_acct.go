package sv

// Keypair is a helper for organizing data about system accounts.
type Keypair struct {
	Public  string
	Private string
}

// SysAcct is a helper for organizing data about system accounts
//
// While the canonical data is in the package-root consts, it is
// useful in i.e. a generation context to be able to package up all the data
// about a particular system account. That's what this is for.
//
// All data in this struct is comprised of strings which name system variables
// which exist either in the SystemStore or in the associated data.
type SysAcct struct {
	Name       string
	Ownership  Keypair
	Validation Keypair
}
