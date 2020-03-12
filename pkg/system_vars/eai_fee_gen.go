package sv

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/ndau/ndaumath/pkg/address"
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *EAIFee) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Fee"
	o = append(o, 0x82, 0xa3, 0x46, 0x65, 0x65)
	o, err = z.Fee.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Fee")
		return
	}
	// string "To"
	o = append(o, 0xa2, 0x54, 0x6f)
	if z.To == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.To.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "To")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EAIFee) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Fee":
			bts, err = z.Fee.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Fee")
				return
			}
		case "To":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.To = nil
			} else {
				if z.To == nil {
					z.To = new(address.Address)
				}
				bts, err = z.To.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "To")
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *EAIFee) Msgsize() (s int) {
	s = 1 + 4 + z.Fee.Msgsize() + 3
	if z.To == nil {
		s += msgp.NilSize
	} else {
		s += z.To.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EAIFeeTable) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		// map header, size 2
		// string "Fee"
		o = append(o, 0x82, 0xa3, 0x46, 0x65, 0x65)
		o, err = z[za0001].Fee.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, za0001, "Fee")
			return
		}
		// string "To"
		o = append(o, 0xa2, 0x54, 0x6f)
		if z[za0001].To == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].To.MarshalMsg(o)
			if err != nil {
				err = msgp.WrapError(err, za0001, "To")
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EAIFeeTable) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(EAIFeeTable, zb0002)
	}
	for zb0001 := range *z {
		var field []byte
		_ = field
		var zb0003 uint32
		zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, zb0001)
			return
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			switch msgp.UnsafeString(field) {
			case "Fee":
				bts, err = (*z)[zb0001].Fee.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, zb0001, "Fee")
					return
				}
			case "To":
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z)[zb0001].To = nil
				} else {
					if (*z)[zb0001].To == nil {
						(*z)[zb0001].To = new(address.Address)
					}
					bts, err = (*z)[zb0001].To.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, zb0001, "To")
						return
					}
				}
			default:
				bts, err = msgp.Skip(bts)
				if err != nil {
					err = msgp.WrapError(err, zb0001)
					return
				}
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EAIFeeTable) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0004 := range z {
		s += 1 + 4 + z[zb0004].Fee.Msgsize() + 3
		if z[zb0004].To == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0004].To.Msgsize()
		}
	}
	return
}
