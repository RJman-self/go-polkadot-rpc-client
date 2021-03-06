package expand

import (
	"fmt"
	"github.com/huandu/xstrings"
	"github.com/rjman-self/go-polkadot-rpc-client/expand/chainx/xevents"
	"github.com/rjman-self/go-polkadot-rpc-client/utils"
	"github.com/rjmand/go-substrate-rpc-client/v2/scale"
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
)

const ChinaX string = "chainx"

type xTransferCall struct {
	Value interface{}
}

func (t *xTransferCall) Decode(decoder scale.Decoder) error {
	//1. 先获取callidx
	b := make([]byte, 2)
	err := decoder.Read(b)
	if err != nil {
		return fmt.Errorf("deode transfer call: read callIdx bytes error: %v", err)
	}
	callIdx := xstrings.RightJustify(utils.IntToHex(b[0]), 2, "0") + xstrings.RightJustify(utils.IntToHex(b[1]), 2, "0")
	result := map[string]interface{}{
		"call_index": callIdx,
	}
	var param []ExtrinsicParam

	// 0 ---> 	Address
	var address MultiAddress
	err = decoder.Decode(&address)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.Address error: %v", err)
	}
	param = append(param,
		ExtrinsicParam{
			Name:     "dest",
			Type:     "Address",
			Value:    utils.BytesToHex(address.AccountId[:]),
			ValueRaw: utils.BytesToHex(address.AccountId[:]),
		})

	// 1 ---> 	AssetId
	var optionId types.UCompact
	err = decoder.Decode(&optionId)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.AssetId error: %v", err)
	}
	assetId := types.U32(utils.UCompactToBigInt(optionId).Uint64())
	param = append(param,
		ExtrinsicParam{
			Name:  "id",
			Type:  "Compact<AssetId>",
			Value: assetId,
		})

	// 2 ----> Compact<Balance>
	var balance types.UCompact

	err = decoder.Decode(&balance)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.Compact<Balance> error: %v", err)
	}
	amount := utils.UCompactToBigInt(balance).Int64()
	param = append(param,
		ExtrinsicParam{
			Name:  "value",
			Type:  "Compact<Balance>",
			Value: amount,
		})
	result["call_args"] = param
	t.Value = result
	return nil
}

type xTransferOpaqueCall struct {
	Value interface{}
}

func (t *xTransferOpaqueCall) Decode(decoder scale.Decoder) error {
	//1. 先获取callidx
	b := make([]byte, 2)
	err := decoder.Read(b)
	if err != nil {
		return fmt.Errorf("deode transfer call: read callIdx bytes error: %v", err)
	}
	callIdx := xstrings.RightJustify(utils.IntToHex(b[0]), 2, "0") + xstrings.RightJustify(utils.IntToHex(b[1]), 2, "0")
	result := map[string]interface{}{
		"call_index": callIdx,
	}
	var param []ExtrinsicParam
	// 0 ---> 	Address
	var address types.AccountID
	err = decoder.Decode(&address)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.Address error: %v", err)
	}
	param = append(param,
		ExtrinsicParam{
			Name:     "dest",
			Type:     "Address",
			Value:    utils.BytesToHex(address[:]),
			ValueRaw: utils.BytesToHex(address[:]),
		})

	// 1 ---> 	AssetId
	var optionId types.UCompact
	err = decoder.Decode(&optionId)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.AssetId error: %v", err)
	}
	assetId := xevents.AssetId(utils.UCompactToBigInt(optionId).Int64())

	param = append(param,
		ExtrinsicParam{
			Name:  "id",
			Type:  "Compact<AssetId>",
			Value: assetId,
		})

	// 2 ----> Compact<Balance>
	var balance types.UCompact
	err = decoder.Decode(&balance)
	if err != nil {
		return fmt.Errorf("decode call: decode XAssets.transfer.Compact<Balance> error: %v", err)
	}
	amount := utils.UCompactToBigInt(balance).Int64()
	param = append(param,
		ExtrinsicParam{
			Name:  "value",
			Type:  "Compact<Balance>",
			Value: amount,
		})
	result["call_args"] = param
	t.Value = result
	return nil
}
