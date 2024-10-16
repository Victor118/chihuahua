package tokenfactory

import (
	"cosmossdk.io/math"
	"github.com/ChihuahuaChain/chihuahua/x/tokenfactory/keeper"
	"github.com/ChihuahuaChain/chihuahua/x/tokenfactory/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, bankKeeper types.BankKeeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyEndBlocker)
	iter, err := k.ActiveAirdrop.Iterate(ctx, nil)

	if err != nil {

		return err
	}
	defer iter.Close()

	coinsToSend := sdk.NewCoins()

	for ; iter.Valid(); iter.Next() {
		stakeDrop, err := iter.KeyValue()
		if err != nil {

			return err
		}

		if ctx.BlockHeight() >= stakeDrop.Value.StartBlock && ctx.BlockHeight() < stakeDrop.Value.EndBlock {
			var blockNumber math.Int = math.NewInt(stakeDrop.Value.EndBlock - stakeDrop.Value.StartBlock)
			amountToMintPerBlock := stakeDrop.Value.Amount.Amount.Quo(blockNumber)
			coinToMint := sdk.NewCoin(stakeDrop.Value.Amount.Denom, amountToMintPerBlock)

			coinsToSend = coinsToSend.Add(coinToMint)

		} else if ctx.BlockHeight() >= stakeDrop.Value.EndBlock {
			k.ActiveAirdrop.Remove(ctx, stakeDrop.Key)
		}
	}

	bankKeeper.MintCoins(ctx, types.ModuleName, coinsToSend)
	return bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.FeeCollectorName, sdk.NewCoins(coinsToSend...))

}
