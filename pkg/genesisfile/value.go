package genesisfile

// ----- ---- --- -- -
// Copyright 2019 Oneiro NA, Inc. All Rights Reserved.
//
// Licensed under the Apache License 2.0 (the "License").  You may not use
// this file except in compliance with the License.  You can obtain a copy
// in the file LICENSE in the source distribution or at
// https://www.apache.org/licenses/LICENSE-2.0.txt
// - -- --- ---- -----

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/oneiro-ndev/msgp-well-known-types/wkt"
	"github.com/oneiro-ndev/ndaumath/pkg/eai"
	math "github.com/oneiro-ndev/ndaumath/pkg/types"
	sv "github.com/oneiro-ndev/system_vars/pkg/system_vars"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

// A Value is the innermost type on the chaos chain
//
// GoType must be a string representation of the type of the data. Valid
// values:
//
// - primitives, in which case `Data` must be of the correct type
// - `[]byte`, in which case `Data` must be a string the base64-encoded representation
// - structs which implement `Valuable`
// - lists of structs which implement `Valuable`
type Value struct {
	GoType string      `toml:"go_type"`
	Data   interface{} `toml:"data"`
}

// Unpack gets the appropriate concrete type from this Value.
//
// Value is primarily a text type, dealing with encoded data.
// It is normal to wish to deal with raw data. Unpack produeces this.
//
// Only certain data types are supported:
//
// - bool, int, int64, uint, uint64, string, time.Time: these types are represented
//   directly in the toml file as native data types
// - []byte: represented in the toml file as a base64-encoding; these are unpacked
// - ndaumath/pkg/types.Ndau, Duration, and Timestamp: these types are represented
//   with appropriate toml primitives
// - <type implementing Valuable>: unpacked with its UnmarshalText
// - []<type implementing Valuable>: unpacked with its UnmarshalText
func (v Value) Unpack() (interface{}, error) {
	switch v.GoType {
	case "bool", "int", "int64", "uint", "uint64", "string", "time.Time":
		if getTypeName(v.Data) != v.GoType {
			return nil, fmt.Errorf("invalid encoding of primitive: expect %s; have %s", v.GoType, getTypeName(v.Data))
		}
		return v.Data, nil
	case "[]byte", "[]uint8":
		encoded, ok := v.Data.(string)
		if !ok {
			return nil, fmt.Errorf("invalid encoding of []byte: expect string, have %s", getTypeName(v.Data))
		}
		return base64.StdEncoding.DecodeString(encoded)
	case "math.Duration", "types.Duration", "math.Ndau", "types.Ndau", "math.Timestamp", "types.Timestamp":
		switch v.GoType {
		case "math.Ndau", "types.Ndau":
			return math.Ndau(v.Data.(int64)), nil
		case "math.Duration", "types.Duration":
			return math.ParseDuration(v.Data.(string))
		case "math.Timestamp", "types.Timestamp":
			return math.TimestampFrom(v.Data.(time.Time))
		default:
			return nil, fmt.Errorf("unreachable: nested case in unpack")
		}
	default:
		typename := v.GoType
		isslice := false

		if strings.HasPrefix(v.GoType, "[]") {
			typename = typename[2:]
			isslice = true
		}

		example, exists := valuableRegistry[typename]
		if !exists {
			return nil, fmt.Errorf("unknown Valuable: %s", typename)
		}

		if isslice {
			switch inS := v.Data.(type) {
			case []string:
				sS, vS := sliceOf(example, len(inS))
				var err error
				for idx := range inS {
					err = sS.Index(idx).Interface().(Valuable).UnmarshalText([]byte(inS[idx]))
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("failed to decode %s at idx %d", typename, idx))
					}
					vS[idx] = sS.Index(idx).Interface().(Valuable)
				}
				return vS, err
			case []interface{}:
				sS, vS := sliceOf(example, len(inS))
				var err error
				for idx := range inS {
					in, ok := inS[idx].(string)
					if !ok {
						return nil, fmt.Errorf("%s must be encoded as []string", v.GoType)
					}
					err = sS.Index(idx).Interface().(Valuable).UnmarshalText([]byte(in))
					if err != nil {
						return nil, errors.Wrap(err, fmt.Sprintf("failed to decode %s at idx %d", typename, idx))
					}
					vS[idx] = sS.Index(idx).Interface().(Valuable)
				}
				return vS, err
			default:
				return nil, fmt.Errorf("%s must be encoded as []string", v.GoType)
			}
		}

		in, ok := v.Data.(string)
		if !ok {
			return nil, fmt.Errorf("%s must be encoded as string", v.GoType)
		}
		out := emptyCopy(example)
		err := out.UnmarshalText([]byte(in))
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to decode %s", typename))
		}

		return out, err
	}
}

// PackValue constructs an appropriate Value from this value.
//
// Value is primarily a text type, dealing with encoded data.
// PackValue produeces this encoded data as appropriate.
//
// Only certain data types are supported:
//
// - bool, int, int64, uint, uint64, string, time.Time: these types are represented
//   directly in the toml file as native data types
// - []byte: represented in the toml file as a base64-encoding
// - ndaumath/pkg/types.Ndau, Duration, and Timestamp: these types are represented
//   with appropriate toml primitives
// - <type implementing Valuable>: packed with its MarshalText
// - []<type implementing Valuable>: packed with its MarshalText
func PackValue(v interface{}) (out Value, err error) {
	switch typeName := getTypeName(v); typeName {
	case "bool", "int", "int64", "uint", "uint64", "string", "time.Time":
		out.GoType = typeName
		out.Data = v
	case "[]byte", "[]uint8":
		out.GoType = typeName
		out.Data = base64.StdEncoding.EncodeToString(v.([]byte))
	case "math.Duration", "types.Duration", "math.Ndau", "types.Ndau", "math.Timestamp", "types.Timestamp":
		switch typeName {
		case "math.Ndau", "types.Ndau":
			out.GoType = "math.Ndau"
			out.Data = int64(v.(math.Ndau))
		case "math.Duration", "types.Duration":
			out.GoType = "math.Duration"
			out.Data = v.(math.Duration).String()
		case "math.Timestamp", "types.Timestamp":
			out.GoType = "math.Timestamp"
			out.Data = v.(math.Timestamp).AsTime()
		}
	default:
		if !implementsValuable(v) {
			err = fmt.Errorf(
				"cannot pack %s: %s must implement Valuable",
				typeName,
				bareTypeOf(v, true).Name(),
			)
			return
		}

		out.GoType = typeName

		isslice := false
		if strings.HasPrefix(typeName, "[]") {
			typeName = typeName[2:]
			isslice = true
		}

		if _, ok := valuableRegistry[typeName]; !ok {
			err = fmt.Errorf(
				"%s not in valuable registry; consider calling RegisterValuable in an init func",
				typeName,
			)
			return
		}

		if isslice {
			var vals []Valuable
			vals, err = toValuableSlice(v)
			if err != nil {
				return
			}

			data := make([]string, len(vals))
			for idx := range vals {
				var bytes []byte
				bytes, err = vals[idx].(Valuable).MarshalText()
				err = errors.Wrap(err, fmt.Sprintf("failed to encode %s at idx %d", typeName, idx))
				if err != nil {
					return
				}
				data[idx] = string(bytes)
			}
			out.Data = data
		} else {
			var bytes []byte
			bytes, err = pointerize(v).(Valuable).MarshalText()
			err = errors.Wrap(err, fmt.Sprintf("failed to encode %s", typeName))
			if err != nil {
				return
			}
			out.Data = string(bytes)
		}
	}
	return
}

// IntoBytes converts this value into a []byte representation.
//
// Only certain data types are supported:
//
// - bool, int, int64, uint, uint64, []byte, string: converted using appropriate wkt encodings
// - time.Time: written directly in msgPack format
// - <type implementing msgp.Marshaler>: packed with its MarshalMsg
func (v Value) IntoBytes() ([]byte, error) {
	val, err := v.Unpack()
	if err != nil {
		return nil, err
	}
	switch data := val.(type) {
	case bool:
		return wkt.Bool(data).MarshalMsg(nil)
	case int:
		return wkt.Int(data).MarshalMsg(nil)
	case int64:
		return wkt.Int64(data).MarshalMsg(nil)
	case uint:
		return wkt.Uint(data).MarshalMsg(nil)
	case uint64:
		return wkt.Uint64(data).MarshalMsg(nil)
	case time.Time:
		return msgp.AppendTime(nil, data), nil
	case []byte:
		return wkt.Bytes(data).MarshalMsg(nil)
	case string:
		return wkt.String(data).MarshalMsg(nil)
	case msgp.Marshaler:
		return data.MarshalMsg(nil)
	}

	// we might run into trouble because of our type exceptions: we've defined
	// msgp.Marshaler for certain lists, but in order to represent them in toml
	// _as_ a list, we've unpacked them into their contained types.
	//
	// last chance, then: can we undo the exception and get a marshaler?
	switch v.GoType {
	case "[]eai.RTRow":
		vals := val.([]Valuable)
		rows := make([]eai.RTRow, len(vals))
		for i := range vals {
			rows[i] = *vals[i].(*eai.RTRow)
		}
		return eai.RateTable(rows).MarshalMsg(nil)
	case "[]sv.EAIFee":
		vals := val.([]Valuable)
		rows := make([]sv.EAIFee, len(vals))
		for i := range vals {
			rows[i] = *vals[i].(*sv.EAIFee)
		}
		return sv.EAIFeeTable(rows).MarshalMsg(nil)
	}

	return nil, fmt.Errorf("unsupported type %T: must implement msgp.Marshaler", v.Data)
}
