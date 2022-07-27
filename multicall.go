package multicall

import (
	"context"
	"github.com/ApeX-Protocol/multicall-go/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

type multiCall struct {
	explorerRpc  string
	multiAddress string
	abi          string
}

type ViewCall struct {
	id        string
	target    string
	method    string
	arguments []interface{}
}

func (m *multiCall) CallTargets(ctx context.Context, calls []*ViewCall) (map[string][]interface{}, error) {
	client, err := ethclient.Dial(m.explorerRpc)
	if err != nil {
		return nil, err
	}

	var multiCallList []contract.Multicall2Call
	for _, call := range calls {
		targetAbi, err := abi.JSON(strings.NewReader(m.abi))
		if err != nil {
			return nil, err
		}

		data, err := targetAbi.Pack(call.method, call.arguments)
		if err != nil {
			return nil, err
		}

		multi2Call := contract.Multicall2Call{
			Target:   common.HexToAddress(call.target),
			CallData: data,
		}

		multiCallList = append(multiCallList, multi2Call)
	}

	multi, err := contract.NewFreeMulticall2(common.HexToAddress(m.multiAddress), client)
	if err != nil {
		return nil, err
	}

	multiRes, err := multi.TryAggregate(&bind.CallOpts{
		Context: ctx,
	}, true, multiCallList)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]interface{})
	for i, call := range calls {
		targetAbi, err := abi.JSON(strings.NewReader(m.abi))
		if err != nil {
			return nil, err
		}

		o, err := targetAbi.Unpack(call.target, multiRes[i].ReturnData)
		if err != nil {
			return nil, err
		}
		res[call.id] = o
	}
	return res, nil
}
