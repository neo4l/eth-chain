package chain

import (
	"strings"

	"github.com/neo4l/x/tool"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeployContract(url, privateKey, gasLimit, gasPrice, byteCode string) (string, string, error) {
	address := PrivateKeyToAddress(privateKey)
	nonce, err := GetTransactionCount(url, address)
	if err != nil {
		return "", "", err
	}
	txObj := NewTxObj(nonce, "", "", gasLimit, gasPrice, byteCode)

	signedData, err := txObj.SignedData(privateKey)
	if err != nil {
		return "", "", err
	}

	txHash, err := SendRawTransaction(url, "0x"+signedData)
	if err != nil {
		return "", "", err
	}

	contractAddress := CreateContractAddress(address, nonce)
	return txHash, contractAddress, nil
}

func CallWithBlock(url, jsonAbi, contractAddress, defaultBlock, functionName string, args ...interface{}) (interface{}, error) {
	abiObj, err := abi.JSON(strings.NewReader(jsonAbi))
	if err != nil {
		return "", err
	}
	dataByte, err := abiObj.Pack(functionName, args...)
	if err != nil {
		return "", err
	}

	txObj := NewTxObj("", contractAddress, "", "", "", "0x"+common.Bytes2Hex(dataByte))

	txObj.GasLimit = ""
	txObj.GasPrice = ""
	reply, err := Call(url, txObj, defaultBlock)
	return reply, err
}

// func SignCall(url, privateKey, jsonAbi, contractAddress, functionName string, args ...interface{}) (string, error) {

// 	address := PrivateKeyToAddress(privateKey)
// 	nonce, err := GetTransactionCount(url, address)
// 	if err != nil {
// 		return "", err
// 	}
// 	return SignCallWithNonce(privateKey, nonce, conf.Default_GasLimit, conf.Default_GasPrice, jsonAbi, contractAddress, functionName, args...)
// }

func SignCallWithNonce(privateKey, nonce, gasLimit, gasPrice, jsonAbi, contractAddress, functionName string, args ...interface{}) (string, error) {

	abiObj, err := abi.JSON(strings.NewReader(jsonAbi))
	if err != nil {
		return "", err
	}
	dataByte, err := abiObj.Pack(functionName, args...)
	if err != nil {
		return "", err
	}
	txObj := NewTxObj(nonce, contractAddress, "", gasLimit, gasPrice, common.Bytes2Hex(dataByte))
	txObj.GasLimit = gasLimit
	txObj.GasPrice = gasPrice

	return txObj.SignedData(privateKey)
}

func TxData(jsonAbi, functionName string, args ...interface{}) (string, error) {
	abiObj, err := abi.JSON(strings.NewReader(jsonAbi))
	if err != nil {
		return "", err
	}
	dataByte, err := abiObj.Pack(functionName, args...)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(dataByte), nil
}

func CreateContractAddress(publicKey, nonce string) string {
	return crypto.CreateAddress(common.HexToAddress(publicKey), tool.HexToUintWithoutError(nonce)).Hex()
}

func PrivateKeyToAddress(privateKey string) (address string) {
	key, _ := crypto.HexToECDSA(privateKey)
	return crypto.PubkeyToAddress(key.PublicKey).Hex()
}
