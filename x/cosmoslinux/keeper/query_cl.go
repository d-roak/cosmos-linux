package keeper

import (
	"context"
	"fmt"

	"cosmos-linux/x/cosmoslinux/types"
	"cosmos-linux/x/cosmoslinux/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Cl(goCtx context.Context, req *types.QueryClRequest) (*types.QueryClResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

    id, _ := utils.StartDockerContainer([]string{})

    return &types.QueryClResponse{Text: fmt.Sprintf("%s", id)}, nil
}

