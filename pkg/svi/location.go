package svi

//go:generate msgp
//msgp:tuple Location

// Location is a namespace and key which together identify a unique value on the chaos chain.
//
// Though the keys have human meaning, and are likely to be strings, we still
// represent them with byte slices because there's nothing theoretically
// prohibiting someone from using a jpeg of a kitten as they key to a system
// variable.
type Location struct {
	Namespace []byte
	Key       []byte
}

// NewLocation constructs a Location from a namespace and a key
func NewLocation(ns []byte, key string) Location {
	return Location{
		Namespace: ns,
		Key:       []byte(key),
	}
}
