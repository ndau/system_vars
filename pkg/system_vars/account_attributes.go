package sv

//go:generate msgp -io=0

// AccountAttributesName is the name of the AccountAttributes system variable
//
// This sytem variable stores a map[string]map[string]struct{} which maps account addresses to
// a map of attributes associated with that account. For example, if AccountAttributes[addr]["x"]
// exists, then the account with address 'addr' is an exchange account.
const AccountAttributesName = "AccountAttributes"

// Available account attributes.
const (
	AccountAttributeExchange string = "x"
)

// AccountAttributes is a list of EAI fees and their destinations
type AccountAttributes map[string]map[string]struct{}

// Zeroize implements validatable
func (aa *AccountAttributes) Zeroize() {
	*aa = make(AccountAttributes)
}

func init() {
	RegisterTypeValidator(AccountAttributesName, &AccountAttributes{})
}
