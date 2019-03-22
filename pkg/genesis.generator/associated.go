package generator

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// DefaultAssociated returns the default path to the associated data
func DefaultAssociated(ndauhome string) string {
	return filepath.Join(ndauhome, "chaos", "associated.toml")
}

// Associated tracks associated data which goes with the mocks.
//
// In particular, it's used for tests. For example, we mock up some
// public/private keypairs for the ReleaseFromEndowment transaction.
// The public halves of those keys are written into the mock file,
// but the private halves are communicated to the test suite by means
// of the Associated struct.
type Associated map[string]interface{}

// Dump writes this associated data to a file
func (a Associated) Dump(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := toml.NewEncoder(bufio.NewWriter(f))
	return encoder.Encode(a)
}
