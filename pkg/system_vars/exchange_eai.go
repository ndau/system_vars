package sv

// ExchangeEAIScriptName is the name of the ExchangeEAIScript system variable
//
// This sytem variable stores the chaincode script for calculating EAI rate for exchange accounts.
const ExchangeEAIScriptName = "ExchangeEAIScript"

// In case the system variable isn't present on the blockchain, we use 2% as the default.
// See commands/cmd/chasm/examples/two_percent.chbin for where this constant comes from.
const ExchangeEAIScriptDefault = "oAAlAMgXqASI"
