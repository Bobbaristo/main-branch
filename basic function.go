// here we are

package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

// SetAppleIAPOrderCount set the total number of appleIAPOrder
func (k Keeper) SetAppleIAPOrderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AppleInAppPurchaseOrderCountKey))
	byteKey := types.KeyPrefix(types.AppleInAppPurchaseOrderCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

	// Update AppleIAPOrder count
	k.SetAppleIAPOrderCount(ctx, count+1)
