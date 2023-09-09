package cli

import (
	"strconv"

	"cosmos-linux/x/cosmoslinux/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOutput() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output [machineId]",
		Short: "Query output",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            argsMachineId := string(args[0])

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOutputRequest{
                MachineId: argsMachineId,
            }

			res, err := queryClient.Output(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
