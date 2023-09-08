package keeper

import (
	"context"
	"fmt"

	"cosmos-linux/x/cosmoslinux/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	docker_types "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (k Keeper) Cl(goCtx context.Context, req *types.QueryClRequest) (*types.QueryClResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
    defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), docker_types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

    return &types.QueryClResponse{Text: fmt.Sprintf("Hello %s", containers[0].Image)}, nil
}
