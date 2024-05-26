package types

const (
	// ModuleName defines the module name
	ModuleName = "liquidity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	PoolKeyPrefix = "Pool/value/"
	NextPoolIDPrefix = "Pool/nextID"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_liquidity"
)

var (
	ParamsKey = []byte("p_liquidity")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
