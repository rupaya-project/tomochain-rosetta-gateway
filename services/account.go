// Copyright (c) 2020 TomoChain

package services

import (
	"context"
	tc "github.com/tomochain/tomochain-rosetta-gateway/tomochain-client"
	"math/big"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

type accountAPIService struct {
	client tc.TomoChainClient
}

// NewAccountAPIService creates a new instance of an AccountAPIService.
func NewAccountAPIService(client tc.TomoChainClient) server.AccountAPIServicer {
	return &accountAPIService{
		client: client,
	}
}

// AccountBalance implements the /account/balance endpoint.
func (s *accountAPIService) AccountBalance(
	ctx context.Context,
	request *types.AccountBalanceRequest,
) (*types.AccountBalanceResponse, *types.Error) {
	terr := ValidateNetworkIdentifier(ctx, s.client, request.NetworkIdentifier)
	if terr != nil {
		return nil, terr
	}
	blockHash, err := s.client.GetBlock(ctx, big.NewInt(*(request.BlockIdentifier.Index)))
	resp, err := s.client.GetAccount(ctx, blockHash, request.AccountIdentifier.Address)
	if err != nil {
		return nil, ErrUnableToGetAccount
	}
	return resp, nil
}