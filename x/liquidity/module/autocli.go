package liquidity

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/Srikarrao1/liquidity/api/liquidity/liquidity"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePool",
					Use:            "create-pool [builder] [asset-a] [asset-b] [initial-amount-a] [initial-amount-b]",
					Short:          "Send a create-pool tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "builder"}, {ProtoField: "assetA"}, {ProtoField: "assetB"}, {ProtoField: "initialAmountA"}, {ProtoField: "initialAmountB"}},
				},
				{
					RpcMethod:      "AddLiquidity",
					Use:            "add-liquidity [pool-id] [amount-a] [amount-b]",
					Short:          "Send a add-liquidity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "poolId"}, {ProtoField: "amountA"}, {ProtoField: "amountB"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
