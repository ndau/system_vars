package generator

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oneiro-ndev/chaincode/pkg/vm"
	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/oneiro-ndev/ndaumath/pkg/constants"
	"github.com/oneiro-ndev/ndaumath/pkg/eai"
	"github.com/oneiro-ndev/ndaumath/pkg/signature"
	math "github.com/oneiro-ndev/ndaumath/pkg/types"
	"github.com/oneiro-ndev/system_vars/pkg/genesisfile"
	sv "github.com/oneiro-ndev/system_vars/pkg/system_vars"
	"github.com/pkg/errors"
)

// GenerateIn makes genesis and associated files in a particular location
//
// If `dir` is blank, these files will be stored in a system-defined
// temporary location, usually `/tmp`. Otherwise they will be stored in the
// specified directory.
//
// We don't keep track of these files or clean them up at any point.
// If they are in a temporary directory, on most OSX and Linux systems,
// they will be cleaned up after three days of disuse. We can get away
// with this because they're small.
func GenerateIn(dir string) (gfilepath, asscpath string, err error) {
	var gfile, asscfile *os.File

	gfile, err = ioutil.TempFile(dir, "genesis.*.toml")
	if err != nil {
		return
	}
	gfilepath = gfile.Name()

	asscfile, err = ioutil.TempFile(dir, "assc.*.toml")
	if err != nil {
		return
	}
	asscpath = asscfile.Name()

	err = Generate(gfilepath, asscpath)
	return
}

// Generate creates a genesisfile and associated data.
//
// Both arguments are paths to the files which should be written.
//
// Both files are written as TOML. Existing data in these files is clobbered.
func Generate(gfilepath, associated string) (err error) {
	var gfile genesisfile.GFile
	var ma Associated
	gfile, ma, err = GenerateData()
	if err != nil {
		return
	}

	err = ma.Dump(associated)
	if err != nil {
		err = errors.Wrap(err, "writing associated file")
		return
	}

	err = gfile.Dump(gfilepath)
	if err != nil {
		err = errors.Wrap(err, "writing genesis file")
		return
	}

	return
}

// GenerateData mocks up some system variables without touching the filesystem
func GenerateData() (gfile genesisfile.GFile, assc Associated, err error) {
	gfile = make(genesisfile.GFile)
	assc = make(Associated)

	// set RFE address
	err = generateSystemAccount(assc, gfile.Set, sv.ReleaseFromEndowment)
	if err != nil {
		return
	}

	// set default rate tables
	err = gfile.Set(sv.UnlockedRateTableName, eai.DefaultUnlockedEAI)
	if err != nil {
		err = errors.Wrap(err, "setting unlocked eai table")
		return
	}
	err = gfile.Set(sv.LockedRateTableName, eai.DefaultLockBonusEAI)
	if err != nil {
		err = errors.Wrap(err, "setting locked rate table")
		return
	}

	// make default recourse duration
	err = gfile.Set(sv.DefaultRecourseDurationName, math.Duration(math.Day*2))
	if err != nil {
		err = errors.Wrap(err, "setting default recourse duration")
		return
	}

	// make default tx fee script
	// this one is very simple: unconditionally returns numeric 0
	err = gfile.Set(sv.TxFeeScriptName, vm.MiniAsm("handler 0 zero enddef").Bytes())
	if err != nil {
		err = errors.Wrap(err, "setting tx fee script")
		return
	}

	// min stake for an account to be active
	err = gfile.Set(sv.MinNodeRegistrationStakeName, math.Ndau(1000*constants.QuantaPerUnit))
	if err != nil {
		err = errors.Wrap(err, "setting min stake")
		return
	}

	// make default node goodness script
	// empty: returns the value on top of the stack
	// as goodness functions have the total stake on top of the stack,
	// that's actually not a terrible default
	err = gfile.Set(sv.NodeGoodnessFuncName, vm.MiniAsm("handler 0 enddef").Bytes())
	if err != nil {
		err = errors.Wrap(err, "setting goodness func")
		return
	}

	// make eai fee table
	var eaiFeeTable sv.EAIFeeTable
	eaiFeeTable, err = makeEAIFeeTable()
	err = gfile.Set(sv.EAIFeeTableName, eaiFeeTable)
	if err != nil {
		err = errors.Wrap(err, "setting eai fee table")
		return
	}

	// set default min duration between node rewards nominations
	err = gfile.Set(sv.MinDurationBetweenNodeRewardNominationsName, math.Duration(1*math.Day))
	if err != nil {
		err = errors.Wrap(err, "setting min duration between nnr txs")
		return
	}

	// set nominate node reward
	err = generateSystemAccount(assc, gfile.Set, sv.NominateNodeReward)
	if err != nil {
		return
	}

	// set node reward nomination timeout
	err = gfile.Set(sv.NodeRewardNominationTimeoutName, math.Duration(30*math.Second))
	if err != nil {
		err = errors.Wrap(err, "setting nnr timeout")
		return
	}

	// set up command validator change
	err = generateSystemAccount(assc, gfile.Set, sv.CommandValidatorChange)
	if err != nil {
		return
	}

	// set up record price
	err = generateSystemAccount(assc, gfile.Set, sv.RecordPrice)
	if err != nil {
		return
	}

	// set up set sysvar
	// set RFE address
	err = generateSystemAccount(assc, gfile.Set, sv.SetSysvar)
	if err != nil {
		return
	}

	// set up BPC rules account
	err = generateSystemAccount(assc, gfile.Set, sv.BPCRulesAccount)
	if err != nil {
		return
	}

	// set up node rules account
	err = generateSystemAccount(assc, gfile.Set, sv.NodeRulesAccount)
	if err != nil {
		return
	}

	// set up dispute rules account
	err = generateSystemAccount(assc, gfile.Set, sv.DisputeRulesAccount)
	if err != nil {
		return
	}

	// set up changeschema account
	err = generateSystemAccount(assc, gfile.Set, sv.ChangeSchema)
	if err != nil {
		return
	}

	// set up recordpricenav account
	err = generateSystemAccount(assc, gfile.Set, sv.RecordEndowmentNAV)
	if err != nil {
		return
	}

	// set up node rules account
	err = generateSystemAccount(assc, gfile.Set, sv.NodeRulesAccount)
	if err != nil {
		return
	}

	// set up ExchangeEAIScript
	//
	// We want a default of 20000000000 (2%). Miniasm requires the raw bytes,
	// little-endian:
	//     >>> struct.pack('<q', 20000000000)
	//     b'\x00\xc8\x17\xa8\x04\x00\x00\x00'
	// We can therefore trim the final three digits and use push5
	err = gfile.Set(sv.ExchangeEAIScriptName, vm.MiniAsm("handler 0 push5 00 c8 17 a8 04 enddef").Bytes())
	if err != nil {
		err = errors.Wrap(err, "setting exchange EAI script")
		return
	}

	// set up SIBScript
	//
	// See https://github.com/oneiro-ndev/chaincode_scripts/blob/69dbb74d8471c03f6ca9cd5e0f95192f42189cef/src/sib/sib.chasm
	var script []byte
	script, err = base64.StdEncoding.DecodeString("oAAmAJxpMN0AJgAQpdToAEYFDQLAiiAQjwUPAkEPAkEmABCl1OgACUYlAIhSanQmABCl1OgARgUlAIhSanTEiiUAiFJqdBCPiA==")
	if err != nil {
		err = errors.Wrap(err, "decoding SIB script")
		return
	}
	err = gfile.Set(sv.SIBScriptName, script)
	if err != nil {
		err = errors.Wrap(err, "setting SIB script")
		return
	}

	// set up EAI overtime default
	err = gfile.Set(sv.EAIOvertime, math.Duration(30*math.Day))
	if err != nil {
		err = errors.Wrap(err, "setting EAI Overtime duration")
		return
	}

	return
}

func generateSystemAccount(
	assc Associated,
	sets func(key string, val interface{}) error,
	acct sv.SysAcct,
) (err error) {
	// generate ownership keys
	assc[acct.Ownership.Public], assc[acct.Ownership.Private], err = signature.Generate(signature.Ed25519, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("generating %s ownership keys", acct.Name))
	}
	// generate validation keys
	assc[acct.Validation.Public], assc[acct.Validation.Private], err = signature.Generate(signature.Ed25519, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("generating %s validation keys", acct.Name))
	}
	// now generate and set the address
	ownership := assc[acct.Ownership.Public].(signature.PublicKey)
	var addr address.Address
	addr, err = address.Generate(address.KindNdau, ownership.KeyBytes())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("generating %s address", acct.Name))
	}
	err = sets(acct.Address, addr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("setting %s address", acct.Name))
	}
	assc[acct.Address] = addr.String()

	return
}
