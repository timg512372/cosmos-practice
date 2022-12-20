package keeper

import (
	"context"

	"blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    Msg.Body,
	}

	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{Id: id}, nil
}
