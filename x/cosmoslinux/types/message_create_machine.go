package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMachine = "create_machine"

var _ sdk.Msg = &MsgCreateMachine{}

func NewMsgCreateMachine(creator string) *MsgCreateMachine {
	return &MsgCreateMachine{
		Creator: creator,
	}
}

func (msg *MsgCreateMachine) Route() string {
	return RouterKey
}

func (msg *MsgCreateMachine) Type() string {
	return TypeMsgCreateMachine
}

func (msg *MsgCreateMachine) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMachine) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMachine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
