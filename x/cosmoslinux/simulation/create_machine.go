package simulation

import (
	"math/rand"

	"cosmos-linux/x/cosmoslinux/keeper"
	"cosmos-linux/x/cosmoslinux/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateMachine(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateMachine{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateMachine simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateMachine simulation not implemented"), nil, nil
	}
}
