package example

import (
	"context"
	"github.com/ApeX-Protocol/multicall-go"
	"github.com/ApeX-Protocol/multicall-go/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func GetBalance(ctx context.Context, address string) (*big.Int, error) {
	mc := multicall.NewMultiCall(multicall.ArbitrumMainnetRpc,
		multicall.ArbitrumMainnetMultiCallAddress)

	calls := []*multicall.ViewCall{
		{"key-1", contract.ERC20MetaData.ABI,
			"0x52A8845DF664D76C69d2EEa607CD793565aF42B8",
			"balanceOf",
			[]interface{}{common.HexToAddress(address)}},
	}

	result, err := mc.CallTargets(ctx, calls)
	if err != nil {
		return nil, err
	}

	gotVal, ok := result["key-1"]
	if !ok {
		return nil, err
	}

	balance := gotVal[0].(*big.Int)
	return balance, nil
}
