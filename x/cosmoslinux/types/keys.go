package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmoslinux"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmoslinux"

	MachineKey = "Machine-Value-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
