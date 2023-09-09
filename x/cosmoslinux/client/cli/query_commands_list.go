package cli

import (
	"strconv"

	"cosmos-linux/x/cosmoslinux/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCommandsList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commands-list [machineId]",
		Short: "Query commands_list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            argsMachineId := string(args[0])

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCommandsListRequest{
                MachineId: argsMachineId,
            }

			res, err := queryClient.CommandsList(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
