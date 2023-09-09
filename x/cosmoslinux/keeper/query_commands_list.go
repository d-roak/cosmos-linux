package keeper

import (
	"context"

	"cosmos-linux/x/cosmoslinux/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CommandsList(goCtx context.Context, req *types.QueryCommandsListRequest) (*types.QueryCommandsListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

    if req.MachineId == "" {
        return nil, status.Error(codes.InvalidArgument, "invalid machine id")
    }

	return &types.QueryCommandsListResponse{}, nil
}
