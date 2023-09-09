package keeper

import (
	"context"

	"cosmos-linux/x/cosmoslinux/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateMachine(goCtx context.Context, msg *types.MsgCreateMachine) (*types.MsgCreateMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMachineResponse{}, nil
}
