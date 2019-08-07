package sv_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/oneiro-ndev/chaincode/pkg/vm"
	"github.com/oneiro-ndev/ndaumath/pkg/address"
	"github.com/oneiro-ndev/ndaumath/pkg/constants"
	"github.com/oneiro-ndev/ndaumath/pkg/eai"
	math "github.com/oneiro-ndev/ndaumath/pkg/types"
	sv "github.com/oneiro-ndev/system_vars/pkg/system_vars"
	"github.com/stretchr/testify/require"
)

func ndau(t *testing.T) []byte {
	value := math.Ndau(rand.Int63())
	m, err := value.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func duration(t *testing.T) []byte {
	value := math.Duration(rand.Int63())
	m, err := value.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func addr(t *testing.T) []byte {
	data := make([]byte, address.MinDataLength*2)
	_, err := rand.Read(data)
	require.NoError(t, err)
	addr, err := address.Generate(address.KindUser, data)
	require.NoError(t, err)
	m, err := addr.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func chaincode(t *testing.T) []byte {
	cc := vm.MiniAsm("handler 0 zero enddef")
	m, err := cc.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func rateTable(t *testing.T) []byte {
	m, err := eai.DefaultUnlockedEAI.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func accountAttributes(t *testing.T) []byte {
	qty := 1 + rand.Intn(10)
	aa := make(sv.AccountAttributes)
	for i := 0; i < qty; i++ {
		data := make([]byte, address.MinDataLength*2)
		_, err := rand.Read(data)
		require.NoError(t, err)
		addr, err := address.Generate(address.KindUser, data)
		require.NoError(t, err)
		aa[addr.String()] = map[string]struct{}{sv.AccountAttributeExchange: struct{}{}}
	}
	m, err := aa.MarshalMsg(nil)
	require.NoError(t, err)
	return m
}

func feeTable(t *testing.T) []byte {
	qty := 1 + rand.Intn(10)
	ft := make(sv.EAIFeeTable, qty)
	for idx := range ft {
		data := make([]byte, address.MinDataLength*2)
		_, err := rand.Read(data)
		require.NoError(t, err)
		addr, err := address.Generate(address.KindUser, data)
		ft[idx] = sv.EAIFee{
			Fee: constants.NapuPerNdau / 100 * math.Ndau(rand.Int63n(constants.NapuPerNdau/10)),
			To:  &addr,
		}
	}
	m, err := ft.MarshalMsg(nil)
	require.NoError(t, err)
	t.Log(ft)
	t.Logf("%x", m)
	return m
}

func reverse(t *testing.T, data []byte) []byte {
	out := make([]byte, len(data))
	for idx := range data {
		out[len(data)-idx-1] = data[idx]
	}
	return out
}

func TestValidators(t *testing.T) {
	cases := []struct {
		name       string
		makeData   func(t *testing.T) []byte
		invalidate func(t *testing.T, data []byte) []byte
	}{
		{sv.MinNodeRegistrationStakeName, ndau, nil},
		{sv.SIBScriptName, chaincode, reverse},
		{sv.RecordPriceAddressName, addr, reverse},
		{sv.ExchangeEAIScriptName, chaincode, reverse},
		{sv.ChangeSchemaAddressName, addr, reverse},
		{sv.DisputeRulesAccountAddressName, addr, reverse},
		{sv.NodeGoodnessFuncName, chaincode, reverse},
		{sv.TxFeeScriptName, chaincode, reverse},
		{sv.EAIOvertime, duration, nil},
		{sv.UnlockedRateTableName, rateTable, reverse},
		{sv.LockedRateTableName, rateTable, reverse},
		{sv.AccountAttributesName, accountAttributes, reverse},
		{sv.RecordEndowmentNAVAddressName, addr, reverse},
		{sv.ReleaseFromEndowmentAddressName, addr, reverse},
		{sv.DefaultRecourseDurationName, duration, nil},
		{sv.SetSysvarAddressName, addr, reverse},
		{sv.CommandValidatorChangeAddressName, addr, reverse},
		{sv.NodeRulesAccountAddressName, addr, reverse},
		{sv.NominateNodeRewardAddressName, addr, reverse},
		{sv.EAIFeeTableName, feeTable, reverse},
		{sv.BPCRulesAccountAddressName, addr, reverse},
	}

	t.Run("Valid", func(t *testing.T) {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				require.True(t, sv.HasValidator(tc.name))
				v := sv.IsValid(tc.name, tc.makeData(t))
				require.NotNil(t, v)
				require.True(t, *v)
			})
		}
	})
	t.Run("Leftovers", func(t *testing.T) {
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				v := sv.IsValid(tc.name, append(tc.makeData(t), 0x00))
				require.NotNil(t, v)
				require.False(t, *v)
			})
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		for _, tc := range cases {
			if tc.invalidate != nil {
				t.Run(tc.name, func(t *testing.T) {
					v := sv.IsValid(tc.name, tc.invalidate(t, tc.makeData(t)))
					require.NotNil(t, v)
					require.False(t, *v)
				})
			}
		}
	})
	t.Run("Unknown", func(t *testing.T) {
		data := make([]byte, 64)
		_, err := rand.Read(data)
		require.NoError(t, err)
		name := fmt.Sprintf("%q", data)[:32]
		require.False(t, sv.HasValidator(name))
		v := sv.IsValid(name, data)
		require.Nil(t, v)
	})
}

func TestAccountAttributesSelfValidation(t *testing.T) {
	aa := make(sv.AccountAttributes)

	data := make([]byte, address.MinDataLength*2)
	_, err := rand.Read(data)
	require.NoError(t, err)
	addr, err := address.Generate(address.KindUser, data)
	require.NoError(t, err)
	aa[sv.AccountAttributeExchange] = map[string]struct{}{addr.String(): struct{}{}}

	m, err := aa.MarshalMsg(nil)
	require.NoError(t, err)

	v := sv.IsValid(sv.AccountAttributesName, m)
	require.NotNil(t, v)
	require.False(t, *v)
}

func TestChaincodeSelfValidation(t *testing.T) {
	// ensure chaincode has the right semantic properties
	cc := vm.MiniAsm("zero one push1 2 ifnz")
	data, err := cc.MarshalMsg(nil)
	require.NoError(t, err)

	v := sv.ValidateChaincode(data)
	require.False(t, v)

	vp := sv.IsValid(sv.SIBScriptName, data)
	require.NotNil(t, vp)
	require.False(t, *vp)
}
