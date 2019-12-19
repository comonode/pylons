package keep

import (
	"errors"

	"github.com/MikeSofaer/pylons/x/pylons/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCookbook sets the cookbook with the name as the key
func (k Keeper) SetCookbook(ctx sdk.Context, cookbook types.Cookbook) error {
	if cookbook.Sender.Empty() {
		return errors.New("SetCookbook: the sender cannot be empty")
	}

	return k.SetObject(ctx, types.TypeCookbook, cookbook.ID, k.CookbookKey, cookbook)
}

// GetCookbook returns cookbook based on UUID
func (k Keeper) GetCookbook(ctx sdk.Context, id string) (types.Cookbook, error) {
	cookbook := types.Cookbook{}
	err := k.GetObject(ctx, types.TypeCookbook, id, k.CookbookKey, &cookbook)
	return cookbook, err
}

// UpdateCookbook is used to update the cookbook using the id
func (k Keeper) UpdateCookbook(ctx sdk.Context, id string, cookbook types.Cookbook) error {
	if cookbook.Sender.Empty() {
		return errors.New("UpdateCookbook: the sender cannot be empty")

	}
	return k.UpdateObject(ctx, types.TypeCookbook, id, k.CookbookKey, cookbook)
}

// GetCookbooksIterator returns an iterator for all the cookbooks
func (k Keeper) GetCookbooksIterator(ctx sdk.Context, sender sdk.AccAddress) sdk.Iterator {
	store := ctx.KVStore(k.CookbookKey)
	return sdk.KVStorePrefixIterator(store, []byte(sender.String()))
}

// DeleteCookbook is used to delete a cookbook based on the id
func (k Keeper) DeleteCookbook(ctx sdk.Context, id string) error {
	return k.DeleteObject(ctx, types.TypeCookbook, id, k.CookbookKey)
}
