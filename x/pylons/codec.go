package pylons

import (
	"github.com/MikeSofaer/pylons/x/pylons/msgs"
	"github.com/MikeSofaer/pylons/x/pylons/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(msgs.MsgGetPylons{}, "pylons/GetPylons", nil)
	cdc.RegisterConcrete(msgs.MsgSendPylons{}, "pylons/SendPylons", nil)
	cdc.RegisterConcrete(msgs.MsgCreateCookbook{}, "pylons/CreateCookbook", nil)
	cdc.RegisterConcrete(msgs.MsgUpdateCookbook{}, "pylons/UpdateCookbook", nil)
	cdc.RegisterConcrete(msgs.MsgCreateRecipe{}, "pylons/CreateRecipe", nil)
	cdc.RegisterConcrete(msgs.MsgUpdateRecipe{}, "pylons/UpdateRecipe", nil)
	cdc.RegisterConcrete(msgs.MsgExecuteRecipe{}, "pylons/ExecuteRecipe", nil)
	cdc.RegisterConcrete(msgs.MsgCheckExecution{}, "pylons/CheckExection", nil)
	cdc.RegisterConcrete(msgs.MsgEnableRecipe{}, "pylons/EnableRecipe", nil)
	cdc.RegisterConcrete(msgs.MsgDisableRecipe{}, "pylons/DisableRecipe", nil)

	cdc.RegisterInterface((*types.WeightedParam)(nil), nil)
}
