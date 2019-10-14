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
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// A genesisfile contains a bunch of system variables encoded in a human-readable format

// this is dumb, but it works:
// we can't take the address of true, so we take the address of this instead
var tru = true

// GFile represents a genesisfile: a human-friendly representation of the system
// variables.
type GFile map[string]Value

// DefaultPath returns the default path at which a chaos genesis file is expected
func DefaultPath(ndauhome string) string {
	return filepath.Join(ndauhome, "ndau", "genesis.toml")
}

// Load returns a system variable set loaded from its file
func Load(path string) (GFile, error) {
	gf := make(map[string]Value)
	_, err := toml.DecodeFile(path, &gf)
	if err != nil {
		return gf, errors.Wrap(err, "loading gfile")
	}

	return gf, nil
}

// IntoSysvars converts a GFile into ndau sysvars
func (gf GFile) IntoSysvars() (map[string][]byte, error) {
	sv := make(map[string][]byte)
	for k, v := range gf {
		b, err := v.IntoBytes()
		if err != nil {
			return nil, errors.Wrap(
				err,
				fmt.Sprintf(
					"Encoding value %s into bytes",
					k,
				),
			)
		}
		sv[k] = b
	}
	return sv, nil
}

// Dump stores the GFile in a file
func (gf GFile) Dump(path string) error {
	// prepare file
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "creating output file")
	}
	defer f.Close()

	// marshal
	return errors.Wrap(toml.NewEncoder(f).Encode(gf), "encoding gfile")
}

// Set puts val into the specified location
func (gf GFile) Set(name string, val interface{}) error {
	if !utf8.ValidString(name) {
		return errors.New("keys must be valid utf-8")
	}

	value, err := PackValue(val)
	if err != nil {
		return err
	}

	if gf == nil {
		gf = make(GFile)
	}
	gf[name] = value

	return nil
}
