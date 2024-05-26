package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, builder string, assetA string, assetB string, initialAmountA int32, initialAmountB int32) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:        creator,
		Builder:        builder,
		AssetA:         assetA,
		AssetB:         assetB,
		InitialAmountA: initialAmountA,
		InitialAmountB: initialAmountB,
	}
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
