package keeper

import (
	"context"
	"cosmos-linux/x/cosmoslinux/types"
	"cosmos-linux/x/cosmoslinux/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

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

func (k Keeper) Output(goCtx context.Context, req *types.QueryOutputRequest) (*types.QueryOutputResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

    output, _ := utils.StartDockerContainer([]string{"ls", "ps"})

    return &types.QueryOutputResponse{Output: output}, nil
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
