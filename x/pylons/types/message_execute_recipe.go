package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgExecuteRecipe{}

func NewMsgExecuteRecipe(creator string, cookbookID string, recipeID string, itemIDs []string) *MsgExecuteRecipe {
	return &MsgExecuteRecipe{
		Creator:    creator,
		CookbookID: cookbookID,
		RecipeID:   recipeID,
		ItemIDs:    itemIDs,
	}
}

func (msg *MsgExecuteRecipe) Route() string {
	return RouterKey
}

func (msg *MsgExecuteRecipe) Type() string {
	return "ExecuteRecipe"
}

func (msg *MsgExecuteRecipe) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExecuteRecipe) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteRecipe) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if err = ValidateID(msg.CookbookID); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if err = ValidateID(msg.RecipeID); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	for _, id := range msg.ItemIDs {
		if err = ValidateNumber(id); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}