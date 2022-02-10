package keeper_test

import (
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/fulldivevr/imversed/app"
	poolstypes "github.com/fulldivevr/imversed/x/pools/types"
)

func genPoolAssets(r *rand.Rand) []poolstypes.PoolAsset {
	denoms := []string{"IBC/0123456789ABCDEF012346789ABCDEF", "IBC/denom56789ABCDEF012346789ABCDEF"}
	assets := []poolstypes.PoolAsset{}
	for _, denom := range denoms {
		amt, _ := simtypes.RandPositiveInt(r, sdk.NewIntWithDecimal(1, 40))
		reserveAmt := sdk.NewCoin(denom, amt)
		weight := sdk.NewInt(r.Int63n(9) + 1)
		assets = append(assets, poolstypes.PoolAsset{Token: reserveAmt, Weight: weight})
	}

	return assets
}

func genPoolParams(r *rand.Rand) poolstypes.PoolParams {
	swapFeeInt := int64(r.Intn(1e5))
	swapFee := sdk.NewDecWithPrec(swapFeeInt, 6)

	exitFeeInt := int64(r.Intn(1e5))
	exitFee := sdk.NewDecWithPrec(exitFeeInt, 6)

	// TODO: Randomly generate LBP params
	return poolstypes.PoolParams{
		SwapFee:                  swapFee,
		ExitFee:                  exitFee,
		SmoothWeightChangeParams: nil,
	}
}

func setupPools(maxNumPoolsToGen int) []poolstypes.PoolI {
	r := rand.New(rand.NewSource(10))
	// setup N pools
	pools := make([]poolstypes.PoolI, 0, maxNumPoolsToGen)
	for i := 0; i < maxNumPoolsToGen; i++ {
		assets := genPoolAssets(r)
		params := genPoolParams(r)
		pool, _ := poolstypes.NewPool(uint64(i), params, assets, "FutureGovernorString", time.Now())
		pools = append(pools, pool)
	}
	return pools
}

func BenchmarkPoolsPoolSerialization(b *testing.B) {
	app := app.Setup(false)
	maxNumPoolsToGen := 5000
	pools := setupPools(maxNumPoolsToGen)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := i % maxNumPoolsToGen
		app.PoolsKeeper.MarshalPool(pools[j])
	}
}

func BenchmarkPoolsPoolDeserialization(b *testing.B) {
	app := app.Setup(false)
	maxNumPoolsToGen := 5000
	pools := setupPools(maxNumPoolsToGen)
	marshals := make([][]byte, 0, maxNumPoolsToGen)
	for i := 0; i < maxNumPoolsToGen; i++ {
		bz, _ := app.PoolsKeeper.MarshalPool(pools[i])
		marshals = append(marshals, bz)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := i % maxNumPoolsToGen
		app.PoolsKeeper.UnmarshalPool(marshals[j])
	}
}
