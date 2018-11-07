package svi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_SetOn(t *testing.T) {
	type args struct {
		name    string
		loc     Location
		current uint64
		on      uint64
	}
	tests := []struct {
		name    string
		m       Map
		args    args
		wantErr bool
	}{
		{
			"can set value on empty map at genesis",
			make(Map),
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 0, 0},
			false,
		},
		{
			"can set value on empty map post genesis",
			make(Map),
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 1, 2},
			false,
		},
		{
			"cannot set value for immediate effect postgenesis",
			make(Map),
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 1, 1},
			true,
		},
		{
			"can set value on nonempty map at genesis",
			Map{"bar": DeferredChange{
				Current:  Location{[]byte("barns"), []byte("bar")},
				Future:   Location{[]byte("barns"), []byte("bar")},
				ChangeOn: 0,
			}},
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 0, 0},
			false,
		},
		{
			"can set value on nonempty map postgenesis",
			Map{"bar": DeferredChange{
				Current:  Location{[]byte("barns"), []byte("bar")},
				Future:   Location{[]byte("barns"), []byte("bar")},
				ChangeOn: 0,
			}},
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 1, 2},
			false,
		},
		{
			"can overwrite value on nonempty map at genesis",
			Map{"foo": DeferredChange{
				Current:  Location{[]byte("barns"), []byte("foo")},
				Future:   Location{[]byte("barns"), []byte("foo")},
				ChangeOn: 0,
			}},
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 0, 0},
			false,
		},
		{
			"can overwrite value on nonempty map postgenesis",
			Map{"foo": DeferredChange{
				Current:  Location{[]byte("barns"), []byte("foo")},
				Future:   Location{[]byte("barns"), []byte("foo")},
				ChangeOn: 0,
			}},
			args{"foo", Location{[]byte("ns"), []byte("foo")}, 1, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.SetOn(tt.args.name, tt.args.loc, tt.args.current, tt.args.on); (err != nil) != tt.wantErr {
				t.Errorf("Map.SetOn() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				loc, err := tt.m.Get(tt.args.name, tt.args.on)
				require.NoError(t, err)
				require.Equal(t, tt.args.loc, loc)
			}
		})
	}
}
