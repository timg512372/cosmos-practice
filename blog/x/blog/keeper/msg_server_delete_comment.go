package keeper

import (
	"blog/x/blog/types"
	"context"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteComment(goCtx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	comment, exist := k.GetComment(ctx, msg.CommentID)
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrID, "Comment doesn't exist")
	}

	if msg.PostID != comment.PostID {
		return nil, sdkerrors.Wrapf(types.ErrID, "Post Blog ID does not exist for comment")
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, comment.Id)
	store.Delete(bz)

	return &types.MsgDeleteCommentResponse{}, nil
}

// message MsgDeleteComment {
//   string creator = 1;
//   uint64 commentID = 2;
//   uint64 postID = 3;
//   uint64 id = 4;
// }
