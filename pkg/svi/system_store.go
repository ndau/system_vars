package svi

import (
	"github.com/tinylib/msgp/msgp"
)

// SystemStore types are stores of system variables.
//
// No restriction is placed on their implementation, so long as they
// can get values from namespaced keys.
type SystemStore interface {
	// GetRaw returns the raw bytes for a given namespace and key.
	//
	// This should normally be avoided in favor of the higher-level Get
	// method, but there are use cases which require this kind of low-level
	// access.
	GetRaw(loc Location) ([]byte, error)
	Get(loc Location, value msgp.Unmarshaler) error
}

// GetFrom gets the requested namespaced key from any SystemStore
func GetFrom(ss SystemStore, loc Location, value msgp.Unmarshaler) error {
	return ss.Get(loc, value)
}

// GetSVI returns the System Variable Indirection map from any SystemStore
func GetSVI(ss SystemStore, loc Location) (Map, error) {
	svi := make(Map)
	err := GetFrom(ss, loc, &svi)
	return svi, err
}
