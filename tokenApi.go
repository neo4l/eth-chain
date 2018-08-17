package chain

import (
	"github.com/neo4l/x/tool"

	"github.com/ethereum/go-ethereum/common"
)

const TokenContractABI = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_amount","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"supply","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"forceTransfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_account","type":"address"}],"name":"unfreeze","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_newAddress","type":"uint256"}],"name":"changeLogicProxy","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_account","type":"address"}],"name":"accountStatus","outputs":[{"name":"_status","type":"uint8"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getInitor","outputs":[{"name":"_proxy","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"unfreezeToken","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[],"name":"freezeToken","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_account","type":"address"}],"name":"freeze","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transferOrigin","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"getProxy","outputs":[{"name":"_proxy","type":"address"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_amounts","type":"uint256"}],"name":"destroy","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_issuer","type":"address"},{"name":"_symbol","type":"bytes32"},{"name":"_id","type":"uint256"},{"name":"_maxSupply","type":"uint256"},{"name":"_precision","type":"uint256"},{"name":"_currentSupply","type":"uint256"},{"name":"_closingTime","type":"uint256"},{"name":"_description","type":"string"},{"name":"_hash","type":"uint256"},{"name":"_coreContract","type":"address"}],"name":"init","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"summary","outputs":[{"name":"_id","type":"uint256"},{"name":"_issuer","type":"address"},{"name":"_symbol","type":"bytes32"},{"name":"_maxSupply","type":"uint256"},{"name":"_precision","type":"uint256"},{"name":"_currentSupply","type":"uint256"},{"name":"_description","type":"string"},{"name":"_registerTime","type":"uint256"},{"name":"_closingTime","type":"uint256"},{"name":"_coreContract","type":"address"},{"name":"_hash","type":"uint256"},{"name":"_status","type":"uint8"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_amounts","type":"uint256"}],"name":"issueMore","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"_allowance","type":"uint256"}],"payable":false,"type":"function"},{"inputs":[],"payable":false,"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_issuer","type":"address"},{"indexed":false,"name":"_symbol","type":"bytes32"},{"indexed":false,"name":"_id","type":"uint256"},{"indexed":false,"name":"_maxSupply","type":"uint256"},{"indexed":false,"name":"_precision","type":"uint256"},{"indexed":false,"name":"_currentSupply","type":"uint256"},{"indexed":false,"name":"_closingTime","type":"uint256"},{"indexed":false,"name":"_description","type":"string"},{"indexed":false,"name":"_hash","type":"uint256"},{"indexed":false,"name":"_coreContract","type":"address"}],"name":"TokenCreate","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_id","type":"uint256"},{"indexed":false,"name":"_from","type":"address"},{"indexed":false,"name":"_to","type":"address"},{"indexed":false,"name":"_amount","type":"uint256"}],"name":"ForceTransfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_issuer","type":"address"},{"indexed":false,"name":"_id","type":"uint256"},{"indexed":false,"name":"_amounts","type":"uint256"}],"name":"IssueMore","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_issuer","type":"address"},{"indexed":false,"name":"_id","type":"uint256"},{"indexed":false,"name":"_amounts","type":"uint256"}],"name":"Destroy","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"from","type":"address"},{"indexed":false,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"owner","type":"address"},{"indexed":false,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_res","type":"uint256[]"}],"name":"Init","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_old","type":"uint256"},{"indexed":false,"name":"_new","type":"uint256"}],"name":"ResetCore","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_old","type":"uint256"},{"indexed":false,"name":"_new","type":"uint256"}],"name":"ResetOwner","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_no","type":"uint256"}],"name":"Alert","type":"event"}]`

func TotalSupply(url, tokenAddress string) (supply string, err error) {

	reply, err := CallWithBlock(url, TokenContractABI, tokenAddress, "latest", "totalSupply")
	//log.Printf("SignData: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	return reply.(string), nil
}

func BalanceOf(url, tokenAddress, accountAddress string) (balance string, err error) {
	return BalanceOfWithBlock(url, tokenAddress, accountAddress, "latest")
}

func BalanceOfWithBlock(url, tokenAddress, accountAddress, defaultBlock string) (balance string, err error) {

	reply, err := CallWithBlock(url, TokenContractABI, tokenAddress, defaultBlock, "balanceOf", common.HexToAddress(accountAddress))
	//log.Printf("SignData: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	return reply.(string), nil
}

func Transfer(url, tokenAddress, to, value, nonce, gasLimit, gasPrice, privateKey string) (hash string, err error) {

	_, signedData, err := SignCallWithNonce(privateKey, nonce, gasLimit, gasPrice, TokenContractABI, tokenAddress, "transfer", common.HexToAddress(to), tool.HexToBigInt(value))
	//log.Printf("signedData: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	txHash, err := SendRawTransaction(url, "0x"+signedData)
	return txHash, err
}

func TransferFrom(url, tokenAddress, from, to, value, nonce, gasLimit, gasPrice, privateKey string) (hash string, err error) {

	_, signedData, err := SignCallWithNonce(privateKey, nonce, gasLimit, gasPrice, TokenContractABI, tokenAddress, "transferFrom", common.HexToAddress(from), common.HexToAddress(to), tool.HexToBigInt(value))
	//log.Printf("Transfer: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	txHash, err := SendRawTransaction(url, "0x"+signedData)
	//t.Logf("SendRawTransaction: %s, %s", txHash, err)

	return txHash, err
}

func Approve(url, tokenAddress, spender, value, nonce, gasLimit, gasPrice, privateKey string) (hash string, err error) {

	_, signedData, err := SignCallWithNonce(privateKey, nonce, gasLimit, gasPrice, TokenContractABI, tokenAddress, "approve", common.HexToAddress(spender), tool.HexToBigInt(value))
	//log.Printf("Transfer: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	txHash, err := SendRawTransaction(url, "0x"+signedData)
	//t.Logf("SendRawTransaction: %s, %s", txHash, err)

	return txHash, err
}

func Allowance(url, tokenAddress, owner, spender string) (amount string, err error) {

	reply, err := CallWithBlock(url, TokenContractABI, tokenAddress, "latest", "allowance", common.HexToAddress(owner), tool.HexToBigInt(spender))
	//log.Printf("SignData: %s, %s", signedData, err)
	if err != nil {
		return "", err
	}
	return reply.(string), nil
}

func ReadContract(url, abi, contractAddress, method string) (reply interface{}, err error) {

	return CallWithBlock(url, abi, contractAddress, "latest", method)
}
