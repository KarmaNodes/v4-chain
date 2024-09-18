package revshare

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/x/revshare/keeper"
	"github.com/dydxprotocol/v4-chain/protocol/x/revshare/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	if err := k.SetMarketMapperRevenueShareParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	k.SetUnconditionalRevShareConfigParams(ctx, genState.UnconditionalRevShareConfig)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetMarketMapperRevenueShareParams(ctx)
	unconditionalRevShareConfig, err := k.GetUnconditionalRevShareConfigParams(ctx)
	if err != nil {
		panic(err)
	}
	genesis.UnconditionalRevShareConfig = unconditionalRevShareConfig
	return genesis
}
