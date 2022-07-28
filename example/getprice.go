package example

import (
	"context"
	"github.com/ApeX-Protocol/multicall-go"
	"github.com/ApeX-Protocol/multicall-go/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func GetPrice(ctx context.Context, pairAddress string) (*big.Int, error) {
	mc := multicall.NewMultiCall(multicall.ArbitrumTestnetRpc,
		multicall.ArbitrumTestnetMultiCallAddress)

	calls := []*multicall.ViewCall{
		{"key-1", contract.PriceOracleMetaData.ABI,
			"0xe7F83e97BC5Fd77CCf4CD25c17Fcb566E7c62781",
			"getMarkPrice",
			[]interface{}{common.HexToAddress(pairAddress)}},
	}

	result, err := mc.CallTargets(ctx, calls)
	if err != nil {
		return nil, err
	}

	gotVal, ok := result["key-1"]
	if !ok {
		return nil, err
	}

	price := gotVal[0].(*big.Int)
	return price, nil
}
