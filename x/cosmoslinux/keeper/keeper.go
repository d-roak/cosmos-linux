package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"cosmos-linux/x/cosmoslinux/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddMachine(ctx sdk.Context, machine types.Machine) string {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MachineKey))

    appendedValue := k.cdc.MustMarshal(&machine)
    store.Set([]byte(machine.Id), appendedValue)

    return machine.Id
}

func (k Keeper) GetMachine(ctx sdk.Context, machineId string) (types.Machine, error) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MachineKey))

    var machine types.Machine
    k.cdc.MustUnmarshal(store.Get([]byte(machineId)), &machine)

    if machine.Id == "" {
        return machine, fmt.Errorf("machine %s does not exist", machineId)
    }

    return machine, nil
}

func (k Keeper) AppendCommand(ctx sdk.Context, machineId string, command string) error {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MachineKey))

    machine, err := k.GetMachine(ctx, machineId)
    if err != nil {
        return err
    }
    machine.Commands = append(machine.Commands, command)

    appendedValue := k.cdc.MustMarshal(&machine)
    store.Set([]byte(machine.Id), appendedValue)

    return nil
}
