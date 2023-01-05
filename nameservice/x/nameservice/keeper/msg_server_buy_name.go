package keeper

import (
	"context"

	"nameservice/x/nameservice/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, isFound := k.GetWhois(ctx, msg.Name)

	minPrice := sdk.Coins{sdk.NewInt64Coin("token", 10)}

	price, _ := sdk.ParseCoinsNormalized(whois.Price)
	bid, _ := sdk.ParseCoinsNormalized(msg.Bid)

	owner, _ := sdk.AccAddressFromBech32(whois.Owner)
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)

	if isFound {

		if price.IsAllGT(bid) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		err := k.bankKeeper.SendCoins(ctx, buyer, owner, bid)
		if err != nil {
			return nil, err
		}
	} else {
		if minPrice.IsAllGT(bid) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid is not high enough")
		}

		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, bid)
		if err != nil {
			return nil, err
		}

	}

	newWhois := types.Whois{
		Index: msg.Name,
		Name:  msg.Name,
		Value: whois.Value,
		Price: bid.String(),
		Owner: buyer.String(),
	}

	k.SetWhois(ctx, newWhois)
	return &types.MsgBuyNameResponse{}, nil
}
