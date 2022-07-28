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

type ViewCall struct {
	Id        string
	TargetAbi string
	Target    string
	Method    string
	Arguments []interface{}
}

type MultiCall interface {
	CallTargets(ctx context.Context, calls []*ViewCall) (map[string][]interface{}, error)
}

type multiCall struct {
	explorerRpc  string
	multiAddress string
}

func NewMultiCall(exploreRpc, multiAddress string) *multiCall {
	return &multiCall{
		explorerRpc:  exploreRpc,
		multiAddress: multiAddress,
	}
}

func (m *multiCall) CallTargets(ctx context.Context, calls []*ViewCall) (map[string][]interface{}, error) {
	client, err := ethclient.Dial(m.explorerRpc)
	if err != nil {
		return nil, err
	}

	var multiCallList []contract.Multicall2Call
	for _, call := range calls {
		targetAbi, err := abi.JSON(strings.NewReader(call.TargetAbi))
		if err != nil {

			return nil, err
		}

		data, err := targetAbi.Pack(call.Method, call.Arguments...)
		if err != nil {
			return nil, err
		}

		multi2Call := contract.Multicall2Call{
			Target:   common.HexToAddress(call.Target),
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
		targetAbi, err := abi.JSON(strings.NewReader(call.TargetAbi))
		if err != nil {
			return nil, err
		}

		o, err := targetAbi.Unpack(call.Method, multiRes[i].ReturnData)
		if err != nil {
			return nil, err
		}
		res[call.Id] = o
	}
	return res, nil
}
