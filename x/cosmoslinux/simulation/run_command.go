package simulation

import (
	"math/rand"

	"cosmos-linux/x/cosmoslinux/keeper"
	"cosmos-linux/x/cosmoslinux/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRunCommand(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRunCommand{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RunCommand simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RunCommand simulation not implemented"), nil, nil
	}
}
