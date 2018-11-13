package svi

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Location) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	err = z.Namespace.DecodeMsg(dc)
	if err != nil {
		return
	}
	err = z.Key.DecodeMsg(dc)
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Location) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = z.Namespace.EncodeMsg(en)
	if err != nil {
		return
	}
	err = z.Key.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Location) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o, err = z.Namespace.MarshalMsg(o)
	if err != nil {
		return
	}
	o, err = z.Key.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Location) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	bts, err = z.Namespace.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	bts, err = z.Key.UnmarshalMsg(bts)
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Location) Msgsize() (s int) {
	s = 1 + z.Namespace.Msgsize() + z.Key.Msgsize()
	return
}
