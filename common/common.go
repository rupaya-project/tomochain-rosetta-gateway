// Copyright (c) 2020 TomoChain

package common

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/types"
	tc "github.com/tomochain/tomochain-rosetta-gateway/tomochain-client"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/common/hexutil"
)

const (
	// tomochain network information
	TomoChainBlockchain       = "tomochain"
	TomoChainMainnetNetWorkId = 88

	// text
	SUCSESS                    = "SUCCESS"
	FAIL                       = "FAIL"
	METADATA_NEW_BALANCE       = "new_balance"
	METADATA_SEQUENCE_NUMBER   = "sequence_number"
	METADATA_ACCOUNT_SEQUENCE  = "account_sequence"
	METADATA_RECENT_BLOCK_HASH = "recent_block_hash"
	METADATA_GAS_LIMIT         = "gas_limit"
	METADATA_GAS_PRICE         = "gas_price"
	METADATA_RECIPIENT         = "recipient"
	METADATA_SENDER            = "sender"
	METADATA_TRANSACTION_TYPE  = "type"
	METADATA_TRANSACTION_VALUE = "value"
	METADATA_TRANSACTION_DATA  = "data"

	// rpc method name
	RPC_METHOD_SEND_SIGNED_TRANSACTION  = "eth_sendRawTransaction"
	RPC_METHOD_GET_PENDING_TRANSACTIONS = "eth_pendingTransactions"
)

type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}

var (
	TomoNativeCoin = &types.Currency{
		Symbol:   "TOMO",
		Decimals: 18,
	}
)

// ValidateNetworkIdentifier validates the network identifier.
func ValidateNetworkIdentifier(ctx context.Context, client tc.TomoChainClient, ni *types.NetworkIdentifier) *types.Error {
	if ni != nil {
		if ni.Blockchain != TomoChainBlockchain {
			return ErrInvalidBlockchain
		}
		if ni.SubNetworkIdentifier != nil {
			return ErrInvalidSubnetwork
		}

		if chainId, err := client.GetChainID(ctx); err != nil || chainId.Uint64() != TomoChainMainnetNetWorkId {
			return ErrInvalidNetwork
		}
	} else {
		return ErrMissingNID
	}
	return nil
}

func SupportedOperationTypes() []string {
	opTyps := make([]string, 0, len(TransactionLogType_name))
	for _, name := range TransactionLogType_name {
		opTyps = append(opTyps, name)
	}
	return opTyps
}

func SupportedConstructionTypes() []string {
	return []string{
		TransactionLogType_NATIVE_TRANSFER.String(),
	}
}

func IsSupportedConstructionType(typ string) bool {
	for _, styp := range SupportedConstructionTypes() {
		if typ == styp {
			return true
		}
	}
	return false
}
