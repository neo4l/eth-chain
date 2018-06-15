package chain

import (
	"encoding/json"

	"github.com/neo4l/x/tool"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type TxObj struct {
	Nonce       string `json:"nonce,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Value       string `json:"value,omitempty"`
	GasLimit    string `json:"gasLimit,omitempty"`
	GasPrice    string `json:"gasPrice,omitempty"`
	Data        string `json:"data,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Blocknumber string `json:"blockNumber,omitempty"`
	BlockTime   string `json:"blockTime,omitempty"`
}

type TxReceipt struct {
	BlockHash         string  `json:"blockHash,omitempty"`
	BlockNumber       string  `json:"blockNumber,omitempty"`
	ContractAddress   string  `json:"contractAddress,omitempty"`
	CumulativeGasUsed string  `json:"cumulativeGasUsed,omitempty"`
	GasUsed           string  `json:"gasUsed,omitempty"`
	Logs              []TxLog `json:"logs,omitempty"`
	LogsBloom         string  `json:"logsBloom,omitempty"`
	TransactionHash   string  `json:"transactionHash,omitempty"`
	TransactionIndex  string  `json:"transactionIndex,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type TxLog struct {
	Address             string   `json:"address,omitempty"`
	BlockHash           string   `json:"blockHash,omitempty"`
	BlockNumber         string   `json:"blockNumber,omitempty"`
	Data                string   `json:"data,omitempty"`
	LogIndex            string   `json:"logIndex,omitempty"`
	TransactionLogIndex string   `json:"transactionLogIndex,omitempty"`
	Topics              []string `json:"topics,omitempty"`
	TransactionHash     string   `json:"transactionHash,omitempty"`
	TransactionIndex    string   `json:"transactionIndex,omitempty"`
	Type                string   `json:"type,omitempty"`
}

func (o *TxReceipt) GetLogData() string {
	if len(o.Logs) < 1 {
		return ""
	}
	for _, log := range o.Logs {
		if log.TransactionLogIndex == "0x0" {
			return log.Data
		}
	}
	return o.Logs[0].Data
}

func (o *TxReceipt) GetLogTopics() []string {
	if len(o.Logs) < 1 {
		return nil
	}
	for _, log := range o.Logs {
		if log.TransactionLogIndex == "0x0" {
			return log.Topics
		}
	}
	return nil
}

func NewTxObj(nonce, to, value, gasLimit, gasPrice, data string) *TxObj {
	return &TxObj{
		Nonce:    nonce,
		To:       to,
		Value:    value,
		GasLimit: gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}
}

func (o *TxObj) ToJson() []byte {
	bytes, err := json.Marshal(o)
	if err != nil {
		return []byte("{ res:0, resMsg: toJson err }")
	} else {
		return bytes
	}
}

func (tx *TxObj) SignedData(privateKey string) (string, error) {

	signedTx, err := tx.Sign(privateKey)
	if err != nil {
		return "", err
	}
	txb, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(txb), nil
}

func DecodesignData(signedData string) (types.Transaction, error) {
	txb := common.Hex2Bytes(signedData)
	var tx types.Transaction
	err := rlp.DecodeBytes(txb, tx)
	return tx, err
}

func (tx *TxObj) Txhash(privateKey string) (string, error) {

	signedTx, err := tx.Sign(privateKey)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (tx *TxObj) Sign(privateKey string) (*types.Transaction, error) {
	key, err := crypto.ToECDSA(common.Hex2Bytes(privateKey))
	if err != nil {
		return nil, err
	}
	var tempTx *types.Transaction
	if tx.To == "" {
		tempTx = types.NewContractCreation(
			tool.HexToUintWithoutError(tx.Nonce),
			tool.HexToBigInt(tx.Value),
			tool.HexToUintWithoutError(tx.GasLimit),
			tool.HexToBigInt(tx.GasPrice),
			common.FromHex(tx.Data),
		)
	} else {
		tempTx = types.NewTransaction(
			tool.HexToUintWithoutError(tx.Nonce),
			common.HexToAddress(tx.To),
			tool.HexToBigInt(tx.Value),
			tool.HexToUintWithoutError(tx.GasLimit),
			tool.HexToBigInt(tx.GasPrice),
			common.FromHex(tx.Data),
		)
	}
	return types.SignTx(tempTx, types.HomesteadSigner{}, key)
}
