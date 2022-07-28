package multicall

import (
	"context"
	"github.com/ApeX-Protocol/multicall-go/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func Test_multiCall_CallTargets(t *testing.T) {
	type fields struct {
		explorerRpc  string
		multiAddress string
	}
	type args struct {
		ctx   context.Context
		calls []*ViewCall
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"arbitrum one testnet", fields{
			explorerRpc:  ArbitrumTestnetRpc,
			multiAddress: ArbitrumTestnetMultiCallAddress,
		}, args{
			context.Background(),
			[]*ViewCall{
				{"key-1", contract.PriceOracleMetaData.ABI,
					"0xe7F83e97BC5Fd77CCf4CD25c17Fcb566E7c62781",
					"getMarkPrice",
					[]interface{}{common.HexToAddress("0xA0b52dBdB5E4B62c8f3555C047440C555773767a")}},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &multiCall{
				explorerRpc:  tt.fields.explorerRpc,
				multiAddress: tt.fields.multiAddress,
			}
			got, err := m.CallTargets(tt.args.ctx, tt.args.calls)
			if err != nil {
				t.Errorf("CallTargets() error = %v", err)
				return
			}
			if got == nil {
				t.Errorf("CallTargets() got = %v", got)
				return
			}
			if gotVal, ok := got["key-1"]; !ok {
				t.Errorf("CallTargets() got = %v", got)
				return
			} else {
				price := gotVal[0].(*big.Int)
				if price.Cmp(big.NewInt(0)) == 0 {
					t.Errorf("CallTargets() got = %v", got)
					return
				}
			}
		})
	}
}
