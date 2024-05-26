package keeper

import (
	"encoding/binary"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/Srikarrao1/liquidity/x/liquidity/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	logger       log.Logger
	storeKey     storetypes.StoreKey

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	storeKey storetypes.StoreKey,
	authority string,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		logger:       logger,
		storeKey:     storeKey,
		authority:    authority,
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	b := k.cdc.MustMarshal(&pool)
	store.Set(GetPoolIDBytes(pool.Id), b)
}

// GetPool retrieves a pool by its ID.
func (k Keeper) GetPool(ctx sdk.Context, id uint64) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKeyPrefix))
	b := store.Get(GetPoolIDBytes(id))

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetPoolIDBytes converts an ID to byte array.
func GetPoolIDBytes(id uint64) []byte {
	return sdk.Uint64ToBigEndian(id)
}

// GetPoolIDFromBytes converts a byte array to an ID.
func GetPoolIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// SetNextPoolID sets the next pool ID in the store.
func (k Keeper) SetNextPoolID(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextPoolIDPrefix))
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	store.Set([]byte{0}, b)
}

// GetNextPoolID retrieves the next pool ID from the store.
func (k Keeper) GetNextPoolID(ctx sdk.Context) (id uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextPoolIDPrefix))
	bz := store.Get([]byte{0})
	if bz == nil {
		return 0, false
	}

	id = binary.BigEndian.Uint64(bz)
	return id, true
}
