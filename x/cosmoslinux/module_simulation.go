package cosmoslinux

import (
	"math/rand"

	"cosmos-linux/testutil/sample"
	cosmoslinuxsimulation "cosmos-linux/x/cosmoslinux/simulation"
	"cosmos-linux/x/cosmoslinux/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cosmoslinuxsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgRunCommand = "op_weight_msg_run_command"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRunCommand int = 100

	opWeightMsgCreateMachine = "op_weight_msg_create_machine"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMachine int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cosmoslinuxGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cosmoslinuxGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRunCommand int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRunCommand, &weightMsgRunCommand, nil,
		func(_ *rand.Rand) {
			weightMsgRunCommand = defaultWeightMsgRunCommand
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRunCommand,
		cosmoslinuxsimulation.SimulateMsgRunCommand(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMachine int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMachine, &weightMsgCreateMachine, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMachine = defaultWeightMsgCreateMachine
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMachine,
		cosmoslinuxsimulation.SimulateMsgCreateMachine(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRunCommand,
			defaultWeightMsgRunCommand,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosmoslinuxsimulation.SimulateMsgRunCommand(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMachine,
			defaultWeightMsgCreateMachine,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				cosmoslinuxsimulation.SimulateMsgCreateMachine(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
