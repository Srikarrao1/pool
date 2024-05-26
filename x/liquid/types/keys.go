package types

const (
	// ModuleName defines the module name
	ModuleName = "liquid"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_liquid"
)

var (
	ParamsKey = []byte("p_liquid")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
