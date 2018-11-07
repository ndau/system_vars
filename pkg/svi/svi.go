package svi

import (
	"encoding"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/tinylib/msgp/msgp"
)

//go:generate msgp

// DeferredChange is an indirection struct.
//
// It helps address the coordination problem: in order to prevent forks,
// all nodes must update their system indirects simultaneously. Otherwise,
// nodes processing the same block may disagree on the indirect, and therefore
// the value, of a given system variable.
//
// Current should always be the current value at the time of the update,
// whether or not that value is stored in the existing "Current" or "Future"
// section from the previous update.
//
// ChangeOn should always be at least 1 more than the current height at the
// time of an update, and best practice will be to increase the buffer,
// because there is no guarantee that a particular transaction will make it
// onto the expected block.
type DeferredChange struct {
	Current  Location
	Future   Location
	ChangeOn uint64
}

// Map is a map of names to deferred changes
//
// Its keys are the string names of system variables.
// Its values are deferred changes. It is a logic error
// to update an Map such that for each updated system variable,
// the updated ChangeOn <= the current height,
// or such that the new value of Current is not equal to the actual
// current value, but it is not possible to actually validate this without
// requiring a custom transaction type for Map updates.
//
// The BPC is encouraged to ensure that it always generates valid Map
// updates, as failure to do so will likely lead to forks.
type Map map[string]DeferredChange

var _ encoding.TextMarshaler = (*Map)(nil)
var _ encoding.TextUnmarshaler = (*Map)(nil)
var _ msgp.Marshaler = (*Map)(nil)
var _ msgp.Unmarshaler = (*Map)(nil)

// MarshalText implements encoding.TextMarshaler
func (m Map) MarshalText() ([]byte, error) {
	bytes, err := m.MarshalMsg(nil)
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(bytes)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (m Map) UnmarshalText(text []byte) error {
	bytes, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return err
	}
	_, err = m.UnmarshalMsg(bytes)
	return err
}

// Marshal this Map to a byte slice
func (m *Map) Marshal() ([]byte, error) {
	return m.MarshalMsg([]byte{})
}

// Unmarshal the byte slice into an Map
func (m *Map) Unmarshal(bytes []byte) error {
	remainder, err := m.UnmarshalMsg(bytes)
	if len(remainder) > 0 {
		return errors.New("Unmarshal produced remainder bytes")
	}
	return err
}

// Get the location of a system variable as of a specififed height
func (m *Map) Get(name string, height uint64) (loc Location, err error) {
	if m == nil {
		err = errors.New("nil Map")
		return
	}
	deferred, hasKey := map[string]DeferredChange(*m)[name]
	if !hasKey {
		err = fmt.Errorf("Key '%s' not present in Map", name)
		return
	}

	if height >= deferred.ChangeOn {
		loc = deferred.Future
	} else {
		loc = deferred.Current
	}

	return
}

// SetOn sets the location of a named system variable to a given namespace and key as of a particular block.
func (m *Map) SetOn(name string, loc Location, current, on uint64) (err error) {
	if on > 0 && on <= current {
		return errors.New("future value must take effect on a block higher than current")
	}
	currentNsk, err := m.Get(name, current)
	if err == nil {
		map[string]DeferredChange(*m)[name] = DeferredChange{
			Current:  currentNsk,
			Future:   loc,
			ChangeOn: on,
		}
	} else {
		_, hasKey := map[string]DeferredChange(*m)[name]
		if !hasKey {
			// error was probably that the key didn't exist
			err = nil
			map[string]DeferredChange(*m)[name] = DeferredChange{
				Current:  loc,
				Future:   loc,
				ChangeOn: on,
			}
		}
	}
	return
}
