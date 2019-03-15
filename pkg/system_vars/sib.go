package sv

// SIBScriptName is the name of the SIBScript system variable
//
// This sytem variable stores the chaincode script for calculating SIB.
const SIBScriptName = "SIBScript"

// SIBScriptDefault exists in case the system variable isn't present on the
// blockchain.
// See https://github.com/oneiro-ndev/chaincode_scripts/blob/e8289c66fd39b0830cbc06066f771d8eafead370/src/sib/sib.chasm
const SIBScriptDefault = "oAAmABCl1OgADwJGBSYAnGkw3QDDiiAQjwUlAIhSanTBiiUAiFJqdBCPJQCIUmp0CSUAiFJqdEElAIhSanRJJgCcaTDdACUAiFJqdEFGQIg="
