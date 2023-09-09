package keeper

import (
	"context"
	"cosmos-linux/x/cosmoslinux/types"
	"crypto/md5"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
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

	commands_hash := md5.Sum([]byte(""))

	machine := types.Machine{
		Creator:  msg.Creator,
		Id:       strings.ReplaceAll(uuid.New().String(), "-", ""),
		State:    fmt.Sprintf("%x", commands_hash),
		Commands: []string{},
	}

	id := k.AddMachine(ctx, machine)
	fmt.Println(id)

	return &types.MsgCreateMachineResponse{MachineId: id}, nil
}

func (k msgServer) RunCommand(goCtx context.Context, msg *types.MsgRunCommand) (*types.MsgRunCommandResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AppendCommand(ctx, msg.MachineId, msg.Command)
	if err != nil {
		return nil, err
	}

	return &types.MsgRunCommandResponse{}, nil
}
