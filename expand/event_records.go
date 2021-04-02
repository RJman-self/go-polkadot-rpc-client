package expand

import (
	"github.com/rjman-self/go-polkadot-rpc-client/expand/bifrost"
	"github.com/rjman-self/go-polkadot-rpc-client/expand/polkadot"
	"github.com/rjmand/go-substrate-rpc-client/v2/types"
	"strings"
)

type IEventRecords interface {
	GetMultisigNewMultisig() []types.EventMultisigNewMultisig
	GetMultisigApproval() []types.EventMultisigApproval
	GetMultisigExecuted() []types.EventMultisigExecuted
	GetMultisigCancelled() []types.EventMultisigCancelled
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetUtilityBatchCompleted() []types.EventUtilityBatchCompleted
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}

/*
扩展： 解析event
*/
func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch strings.ToLower(chainName) {
	case "polkadot":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "kusama":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "westend":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		var events bifrost.BifrostEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events

	}

	return ier, nil
}
