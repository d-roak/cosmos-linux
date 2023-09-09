package keeper

import (
	"context"
	"cosmos-linux/x/cosmoslinux/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateMachine(goCtx context.Context, msg *types.MsgCreateMachine) (*types.MsgCreateMachineResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateMachineResponse{}, nil
}

func (k msgServer) RunCommand(goCtx context.Context, msg *types.MsgRunCommand) (*types.MsgRunCommandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRunCommandResponse{}, nil
}
