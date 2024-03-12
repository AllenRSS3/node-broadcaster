// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethereum

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*logMarshal)(nil)

// MarshalJSON marshals as JSON.
func (l Log) MarshalJSON() ([]byte, error) {
	type Log struct {
		Address          common.Address `json:"address"`
		Topics           []common.Hash  `json:"topics"`
		Data             hexutil.Bytes  `json:"data"`
		BlockNumber      *hexutil.Big   `json:"blockNumber"`
		TransactionHash  common.Hash    `json:"transactionHash"`
		TransactionIndex hexutil.Uint   `json:"transactionIndex"`
		BlockHash        common.Hash    `json:"blockHash"`
		Index            hexutil.Uint   `json:"logIndex"`
		Removed          bool           `json:"removed"`
	}
	var enc Log
	enc.Address = l.Address
	enc.Topics = l.Topics
	enc.Data = l.Data
	enc.BlockNumber = (*hexutil.Big)(l.BlockNumber)
	enc.TransactionHash = l.TransactionHash
	enc.TransactionIndex = hexutil.Uint(l.TransactionIndex)
	enc.BlockHash = l.BlockHash
	enc.Index = hexutil.Uint(l.Index)
	enc.Removed = l.Removed
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (l *Log) UnmarshalJSON(input []byte) error {
	type Log struct {
		Address          *common.Address `json:"address"`
		Topics           []common.Hash   `json:"topics"`
		Data             *hexutil.Bytes  `json:"data"`
		BlockNumber      *hexutil.Big    `json:"blockNumber"`
		TransactionHash  *common.Hash    `json:"transactionHash"`
		TransactionIndex *hexutil.Uint   `json:"transactionIndex"`
		BlockHash        *common.Hash    `json:"blockHash"`
		Index            *hexutil.Uint   `json:"logIndex"`
		Removed          *bool           `json:"removed"`
	}
	var dec Log
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Address != nil {
		l.Address = *dec.Address
	}
	if dec.Topics != nil {
		l.Topics = dec.Topics
	}
	if dec.Data != nil {
		l.Data = *dec.Data
	}
	if dec.BlockNumber != nil {
		l.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.TransactionHash != nil {
		l.TransactionHash = *dec.TransactionHash
	}
	if dec.TransactionIndex != nil {
		l.TransactionIndex = uint(*dec.TransactionIndex)
	}
	if dec.BlockHash != nil {
		l.BlockHash = *dec.BlockHash
	}
	if dec.Index != nil {
		l.Index = uint(*dec.Index)
	}
	if dec.Removed != nil {
		l.Removed = *dec.Removed
	}
	return nil
}