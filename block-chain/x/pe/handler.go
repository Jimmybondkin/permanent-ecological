package pe

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	gtypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/sdk-tutorials/pe/x/pe/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "pe" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgAddCoin:
			return handleMsgAddCoin(ctx, keeper, msg)
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized pe Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle a message to set name
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}

// Handle a message to buy name
func handleMsgAddCoin(ctx sdk.Context, keeper Keeper, msg MsgAddCoin) sdk.Result {
	//check if admin
	cliCtx := context.NewCLIContext()
	resultGenesis, err := cliCtx.Client.Genesis()
	if err != nil {
		return types.ErrGenesisReadError(types.DefaultCodespace).Result()
	}

	appState, err := gtypes.GenesisStateFromGenDoc(cliCtx.Codec, *resultGenesis.Genesis)
	if err != nil {
		return types.ErrGenesisReadError(types.DefaultCodespace).Result()
	}

	genState := gtypes.GetGenesisStateFromAppState(cliCtx.Codec, appState)
	genTxs := make([]sdk.Tx, len(genState.GenTxs))
	for i, tx := range genState.GenTxs {
		err := cliCtx.Codec.UnmarshalJSON(tx, &genTxs[i])
		if err != nil {
			return types.ErrGenesisReadError(types.DefaultCodespace).Result()
		}
	}
	firstTxSigner := genTxs[0].GetMsgs()[0].GetSigners()[0]
	if firstTxSigner.String() != msg.Admin.String() {
		return types.ErrNotTheAdminError(types.DefaultCodespace).Result()
	}
	// Checks if the the bid price is greater than the price paid by the current owner
	_, err = keeper.CoinKeeper.AddCoins(ctx, msg.Admin, msg.Bid) // If so, deduct the Bid amount from the sender
	if err != nil {
		return types.ErrCoinAddError(types.DefaultCodespace).Result()
	}
	return sdk.Result{}
}

// Handle a message to buy name
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) {
		return sdk.ErrInsufficientCoins("Bid not high enough").Result() // If not, throw an error
	}
	if keeper.HasOwner(ctx, msg.Name) {
		err := keeper.CoinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, err := keeper.CoinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) // If so, deduct the Bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg MsgDeleteName) sdk.Result {
	if !keeper.IsNamePresent(ctx, msg.Name) {
		return types.ErrNameDoesNotExist(types.DefaultCodespace).Result()
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}

	keeper.DeleteWhois(ctx, msg.Name)
	return sdk.Result{}
}
