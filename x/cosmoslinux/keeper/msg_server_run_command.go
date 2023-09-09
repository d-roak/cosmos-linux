package keeper

import (
	"context"

	"cosmos-linux/x/cosmoslinux/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RunCommand(goCtx context.Context, msg *types.MsgRunCommand) (*types.MsgRunCommandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRunCommandResponse{}, nil
}
