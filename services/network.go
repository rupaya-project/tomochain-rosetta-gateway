// Copyright (c) 2020 TomoChain

package services

import (
	"context"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/tomochain/tomochain-rosetta-gateway/common"
	tc "github.com/tomochain/tomochain-rosetta-gateway/tomochain-client"
	"github.com/tomochain/tomochain/params"
	"strconv"
)

const (
	latestVersion = "v1.1.0"
)

type networkAPIService struct {
	client tc.TomoChainClient
}

// NewNetworkAPIService creates a new instance of a NetworkAPIService.
func NewNetworkAPIService(client tc.TomoChainClient) server.NetworkAPIServicer {
	return &networkAPIService{
		client: client,
	}
}

// NetworkList implements the /network/list endpoint.
func (s *networkAPIService) NetworkList(
	ctx context.Context,
	request *types.MetadataRequest,
) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{
			{
				Blockchain: common.TomoChainBlockchain,
				Network:    strconv.FormatUint(common.TomoChainMainnetNetWorkId, 10),
			},
		},
	}, nil
}

// NetworkStatus implements the /network/status endpoint.
func (s *networkAPIService) NetworkStatus(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkStatusResponse, *types.Error) {
	terr := common.ValidateNetworkIdentifier(ctx, s.client, request.NetworkIdentifier)
	if terr != nil {
		return nil, terr
	}

	blk, err := s.client.GetBlockByNumber(ctx, nil) // nil means: get latest block
	if err != nil {
		return nil, common.ErrUnableToGetNodeStatus
	}
	genesisblk, err := s.client.GetGenesisBlock(ctx)
	if err != nil {
		return nil, common.ErrUnableToGetNodeStatus
	}

	resp := &types.NetworkStatusResponse{
		CurrentBlockIdentifier: blk.BlockIdentifier,
		CurrentBlockTimestamp:  blk.Timestamp, // ms
		GenesisBlockIdentifier: genesisblk.BlockIdentifier,
	}

	return resp, nil
}

// NetworkOptions implements the /network/options endpoint.
func (s *networkAPIService) NetworkOptions(
	ctx context.Context,
	request *types.NetworkRequest,
) (*types.NetworkOptionsResponse, *types.Error) {
	terr := common.ValidateNetworkIdentifier(ctx, s.client, request.NetworkIdentifier)
	if terr != nil {
		return nil, terr
	}

	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: s.client.GetConfig().Server.RosettaVersion,
			NodeVersion:    params.Version,
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     common.SUCSESS,
					Successful: true,
				},
				{
					Status:     common.FAIL,
					Successful: false,
				},
			},
			OperationTypes: common.SupportedOperationTypes(),
			Errors:         common.ErrorList,
		},
	}, nil
}
