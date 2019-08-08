package sv

// NodeGoodnessFuncName names the chaincode computing the node's goodness functin
const NodeGoodnessFuncName = "NodeGoodnessFunction"

func init() {
	RegisterFuncValidator(NodeGoodnessFuncName, ValidateChaincode)
}
