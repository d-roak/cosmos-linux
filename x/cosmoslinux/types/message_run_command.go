package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRunCommand = "run_command"

var _ sdk.Msg = &MsgRunCommand{}

func NewMsgRunCommand(creator string) *MsgRunCommand {
	return &MsgRunCommand{
		Creator: creator,
	}
}

func (msg *MsgRunCommand) Route() string {
	return RouterKey
}

func (msg *MsgRunCommand) Type() string {
	return TypeMsgRunCommand
}

func (msg *MsgRunCommand) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRunCommand) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRunCommand) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
