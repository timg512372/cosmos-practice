package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(byteKey, appendedValue)

	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Post, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	byteKey := []byte(types.PostKey)
	binary.BigEndian.PutUint64(byteKey, id)

	post := store.Get(byteKey)
	if post == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(post, &val)
	return val, true
}
