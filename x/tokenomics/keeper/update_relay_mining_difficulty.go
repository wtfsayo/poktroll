package keeper

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	proofkeeper "github.com/pokt-network/poktroll/x/proof/keeper"
	prooftypes "github.com/pokt-network/poktroll/x/proof/types"
	"github.com/pokt-network/poktroll/x/tokenomics/types"
)

// TODO_IN_THIS_PR: Prepare future work to decided if these should be
// constants, governance parameters or computed in some other way.
const (
	// Exponential moving average smoothing factor, commonly known as alpha.
	// Large alpha -> more weight on recent data; less smoothing and fast response.
	// Small alpha -> more weight on past data; more smoothing and slow response.
	// Usually, alpha = 2 / (N+1), where N is the number of periods.
	emaSmoothingFactor = float64(0.1)

	// The target number of relays we want the network to mine for a specific
	// service (across all applications & suppliers) per session when claims
	// are aggregated. This number is manually controlled based on how large
	// we want the off-chain trees to grow.
	targetNumRelays = uint64(10e4)
)

// UpdateRelayMiningDifficulty updates the on-chain relay mining difficulty
// based on the amount of on-chain volume.
func (k Keeper) UpdateRelayMiningDifficulty(
	ctx sdk.Context,
	relaysPerServiceMap map[string]uint64,
) error {
	logger := k.Logger().With("method", "UpdateRelayMiningDifficulty")

	for serviceId, numRelays := range relaysPerServiceMap {
		prevDifficulty, found := k.GetRelayMiningDifficulty(ctx, serviceId)
		if !found {
			// If the difficulty is not found, we initialize it with a default.
			prevDifficulty = types.RelayMiningDifficulty{
				ServiceId:    serviceId,
				BlockHeight:  ctx.BlockHeight(),
				NumRelaysEma: 0,
				Difficulty:   defaultDifficultyHash(),
			}
		}

		// TODO_IN_THIS_PR: Should we compute this?
		// N := ctx.BlockHeight() - prevDifficulty.BlockHeight
		// alpha := 2 / (1 + N)
		alpha := emaSmoothingFactor

		// Compute the updated EMA of the number of relays.
		prevRelaysEma := prevDifficulty.NumRelaysEma
		newRelaysEma := computeEma(alpha, prevRelaysEma, numRelays)

		difficultyHash := computeNewDifficultyHash(targetNumRelays, newRelaysEma)
		newDifficulty := types.RelayMiningDifficulty{
			ServiceId:    serviceId,
			BlockHeight:  ctx.BlockHeight(),
			NumRelaysEma: newRelaysEma,
			Difficulty:   difficultyHash,
		}

		k.UpsertRelayMiningDifficulty(ctx, newDifficulty)

		// TODO_IN_THIS_PR: Emit an event for this.
		logger.Info(fmt.Sprintf("Updated relay mining difficulty for service %s from %v to %v", serviceId, prevDifficulty.Difficulty, newDifficulty.Difficulty))

	}
	return nil
}

func computeNewDifficultyHash(targetNumRelays, newRelaysEma uint64) []byte {
	// The target number of relays we want the network to mine is greater than
	// the actual, so we don't need to scale to anything outside the default.
	if targetNumRelays > newRelaysEma {
		return defaultDifficultyHash()
	}

	log2 := func(x float64) float64 {
		return math.Log(x) / math.Ln2
	}

	// Given we are dealing with a bitwise distribution, we are trying to compute
	// the number of leading zeros we need so probabilistically, only T relays
	// are stored in the trees, even though R relays are flowing through the network.
	// 	(0.5)^x = (T/R)
	// 	x = -ln2(T/R)
	numLeadingZeroes := -log2(float64(targetNumRelays) / float64(newRelaysEma))
	numBytes := proofkeeper.SmtSpec.PathHasherSize()
	return leadingZeroBitsToTargetDifficultyHash(int(numLeadingZeroes), numBytes)
}

// computeEma computes the EMA at time t, given the EMA at time t-1, the raw
// data revealed at time t, and the smoothing factor α
// Src: https://en.wikipedia.org/wiki/Exponential_smoothing
func computeEma(alpha float64, prevEma, currValue uint64) uint64 {
	return uint64(alpha*float64(currValue) + (1-alpha)*float64(prevEma))
}

func defaultDifficultyHash() []byte {
	numBytes := proofkeeper.SmtSpec.PathHasherSize()
	numDefaultLeadingZeroBits := int(prooftypes.DefaultMinRelayDifficultyBits)
	return leadingZeroBitsToTargetDifficultyHash(numDefaultLeadingZeroBits, numBytes)
}

// leadingZeroBitsToTargetDifficultyHash generates a slice of bytes with the specified number of leading zero bits
func leadingZeroBitsToTargetDifficultyHash(numLeadingZeroBits int, numBytes int) []byte {
	targetDifficultyHah := make([]byte, numBytes)

	// Set everything to ones initially
	for i := range targetDifficultyHah {
		targetDifficultyHah[i] = 0xff
	}

	// Set full zero bytes
	fullZeroBytes := numLeadingZeroBits / 8
	for i := 0; i < fullZeroBytes; i++ {
		targetDifficultyHah[i] = 0
	}

	// Set remaining bits in the next byte
	remainingZeroBits := numLeadingZeroBits % 8
	if remainingZeroBits > 0 {
		targetDifficultyHah[fullZeroBytes] = byte(0xff >> remainingZeroBits)
	}

	return targetDifficultyHah
}
