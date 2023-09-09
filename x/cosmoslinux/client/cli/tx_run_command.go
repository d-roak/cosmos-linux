package cli

import (
	"strconv"

	"cosmos-linux/x/cosmoslinux/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-command [machineId] [cmd]",
		Short: "Broadcast message run_command",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            argsMachineId := string(args[0])
            argsCmd := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRunCommand(
				clientCtx.GetFromAddress().String(),
                argsMachineId,
                argsCmd,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
