package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	// storetypes "cosmossdk.io/store/types"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	logger       log.Logger
	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	// storeKey storetypes.StoreKey,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
		// storeKey:     storeKey,
		authority: authority,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPool sets a pool in the store.
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))
	appendedValue := k.cdc.MustMarshal(&pool)
	store.Set(GetPoolIDBytes(pool.Id), appendedValue)
}

// GetPool retrieves a pool by its ID.
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (val types.Pool, found bool) {
	
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PoolKeyPrefix))
    b := store.Get(GetPoolIDBytes(id))
    if b == nil {
        return val, false
    }
    k.cdc.MustUnmarshal(b, &val)
    return val, true


}

func GetPoolIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetPoolIDFromBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// SetNextPoolID sets the next pool ID in the store.
func (k Keeper) SetNextPoolID(ctx sdk.Context, id uint64) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.NextPoolIDPrefix)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	store.Set(byteKey, bz)
}

// GetNextPoolID retrieves the next pool ID from the store.
func (k Keeper) GetNextPoolID(ctx sdk.Context) (id uint64, found bool) {
	fmt.Println("csdcds=12")
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.NextPoolIDPrefix)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0, false
	}
	id = binary.BigEndian.Uint64(bz)
	return id, true
}
