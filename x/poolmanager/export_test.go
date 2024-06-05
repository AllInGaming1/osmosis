package poolmanager

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmomath"
	"github.com/osmosis-labs/osmosis/v25/x/poolmanager/types"
)

var IntMaxValue = intMaxValue

func (k Keeper) GetNextPoolIdAndIncrement(ctx sdk.Context) uint64 {
	return k.getNextPoolIdAndIncrement(ctx)
}

// SetPoolRoutesUnsafe sets the given routes to the poolmanager keeper
// to allow routing from a pool type to a certain swap module.
// For example, balancer -> gamm.
// This utility function is only exposed for testing and should not be moved
// outside of the _test.go files.
func (k *Keeper) SetPoolRoutesUnsafe(routes map[types.PoolType]types.PoolModuleI) {
	k.routes = routes
	k.cachedPoolModules = &sync.Map{}
}

// SetPoolModulesUnsafe sets the given modules to the poolmanager keeper.
// This utility function is only exposed for testing and should not be moved
// outside of the _test.go files.
func (k *Keeper) SetPoolModulesUnsafe(poolModules []types.PoolModuleI) {
	k.poolModules = poolModules
	k.cachedPoolModules = &sync.Map{}
}

func (k Keeper) GetAllPoolRoutes(ctx sdk.Context) []types.ModuleRoute {
	return k.getAllPoolRoutes(ctx)
}

func (k Keeper) ValidateCreatedPool(ctx sdk.Context, poolId uint64, pool types.PoolI) error {
	return k.validateCreatedPool(poolId, pool)
}

func (k Keeper) CreateMultihopExpectedSwapOuts(
	ctx sdk.Context,
	route []types.SwapAmountOutRoute,
	tokenOut sdk.Coin,
) ([]osmomath.Int, error) {
	return k.createMultihopExpectedSwapOuts(ctx, route, tokenOut)
}

func (k Keeper) TrackVolume(ctx sdk.Context, poolId uint64, volumeGenerated sdk.Coin) {
	k.trackVolume(ctx, poolId, volumeGenerated)
}

func (k Keeper) CalcAndChargeTakerFee(ctx sdk.Context, tokenIn sdk.Coin, tokenOutDenom string, sender sdk.AccAddress, exactIn bool) (sdk.Coin, error) {
	return k.calcAndChargeTakerFee(ctx, tokenIn, tokenOutDenom, sender, exactIn)
}

func (k Keeper) RouteExactAmountInInternal(
	ctx sdk.Context,
	sender sdk.AccAddress,
	route []types.SwapAmountInRoute,
	tokenIn sdk.Coin,
	tokenOutMinAmount osmomath.Int,
) (tokenOutAmount osmomath.Int, totalTakerFeesToCharge sdk.Coins, err error) {
	return k.routeExactAmountInInternal(ctx, sender, route, tokenIn, tokenOutMinAmount)
}

func (k Keeper) RouteExactAmountOutInternal(
	ctx sdk.Context,
	sender sdk.AccAddress,
	route []types.SwapAmountOutRoute,
	tokenInMaxAmount osmomath.Int,
	tokenOut sdk.Coin,
) (tokenInAmount osmomath.Int, totalTakerFeesToCharge sdk.Coins, err error) {
	return k.routeExactAmountOutInternal(ctx, sender, route, tokenInMaxAmount, tokenOut)
}
