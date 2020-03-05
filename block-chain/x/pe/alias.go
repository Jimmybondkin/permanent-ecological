package pe

import (
	"github.com/cosmos/sdk-tutorials/pe/x/pe/internal/keeper"
	"github.com/cosmos/sdk-tutorials/pe/x/pe/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgAddCoin    = types.NewMsgAddCoin
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
	MsgSetName = types.MsgSetName
	MsgBuyName = types.MsgBuyName
	MsgAddCoin = types.MsgAddCoin
	MsgDeleteName = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames = types.QueryResNames
	Whois = types.Whois
)
