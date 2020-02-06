package genesisfile

// ----- ---- --- -- -
// Copyright 2019, 2020 The Axiom Foundation. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----


import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var valuableType = reflect.TypeOf((*Valuable)(nil)).Elem()

func bareTypeOf(i interface{}, unpackSlice bool) reflect.Type {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr || (unpackSlice && t.Kind() == reflect.Slice) {
		t = t.Elem()
	}
	return t
}

// if input type is a slice of pointers, remove the pointer indirect
func unpointerSlice(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Ptr {
		t = reflect.SliceOf(t.Elem().Elem())
	}
	return t
}

func implementsValuable(i interface{}) bool {
	return reflect.PtrTo(bareTypeOf(i, true)).Implements(valuableType)
}

func getTypeName(i interface{}) string {
	typename := unpointerSlice(bareTypeOf(i, false)).String()
	exc, exists := typeExceptions[typename]
	if exists {
		return exc
	}
	return typename
}

// create an empty copy of an instance
func emptyCopy(original Valuable) Valuable {
	val := reflect.ValueOf(original)
	for val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}
	newThing := reflect.New(val.Type()).Interface().(Valuable)
	return newThing
}

// return two lists: the first is useful for actually getting individual
// values of the correct type for unmarshalling. The second is useful
// for packing the unmarshalled values into, to have a slice of a useful type
func sliceOf(original Valuable, len int) (reflect.Value, []Valuable) {
	sS := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(original)), len, len)
	for i := 0; i < len; i++ {
		sS.Index(i).Set(reflect.ValueOf(emptyCopy(original)))
	}
	vS := reflect.MakeSlice(reflect.SliceOf(valuableType), len, len).Interface().([]Valuable)
	return sS, vS
}

func toValuableSlice(i interface{}) ([]Valuable, error) {
	if !implementsValuable(i) {
		return nil, errors.New("inner type is not Valuable")
	}
	if reflect.TypeOf(i).Kind() != reflect.Slice {
		return nil, errors.New("outer type is not slice")
	}
	iv := reflect.ValueOf(i)
	len := iv.Len()
	vs := reflect.MakeSlice(reflect.SliceOf(valuableType), len, len).Interface().([]Valuable)
	for idx := 0; idx < len; idx++ {
		var ok bool
		vs[idx], ok = pointerize(iv.Index(idx).Interface()).(Valuable)
		if !ok {
			return nil, fmt.Errorf("item at idx %d does not implement Valuable", idx)
		}
	}
	return vs, nil
}

func unpointerTypename(typename string) string {
	switch {
	case strings.HasPrefix(typename, "[]"):
		return "[]" + unpointerTypename(typename[2:])
	case strings.HasPrefix(typename, "*"):
		return unpointerTypename(typename[1:])
	default:
		return typename
	}
}

// if i implements Valuable, return it
// if not, try making a pointer to i: if that implements Valuable, return that
// otherwise return nil
func pointerize(i interface{}) Valuable {
	if cast, ok := i.(Valuable); ok {
		return cast
	}

	rv := reflect.New(reflect.TypeOf(i)) // create pointer
	rv.Elem().Set(reflect.ValueOf(i))    // have the pointer point to the value passed in
	ip := rv.Interface()                 // get the pointer as an empty interface

	if cast, ok := ip.(Valuable); ok {
		return cast
	}
	return nil
}
