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

	machine, err := k.GetMachine(ctx, req.MachineId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid machine id")
	}

	return &types.QueryCommandsListResponse{Commands: machine.Commands}, nil
}

func (k Keeper) Output(goCtx context.Context, req *types.QueryOutputRequest) (*types.QueryOutputResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	machine, err := k.GetMachine(ctx, req.MachineId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid machine id")
	}

	output, _ := utils.ExecuteDockerContainer(machine.Id, machine.Commands)

	return &types.QueryOutputResponse{Output: output}, nil
}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
