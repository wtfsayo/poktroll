package integration_test

import (
	"context"
	"crypto/sha256"
	"testing"

	"github.com/pokt-network/smt"
	"github.com/pokt-network/smt/kvstore/badger"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/cmd/poktrolld/cmd"
	"github.com/pokt-network/poktroll/pkg/crypto/protocol"
	"github.com/pokt-network/poktroll/proto/types/proof"
	"github.com/pokt-network/poktroll/proto/types/session"
	sharedtypes "github.com/pokt-network/poktroll/proto/types/shared"
	"github.com/pokt-network/poktroll/proto/types/tokenomics"
	testutilevents "github.com/pokt-network/poktroll/testutil/events"
	"github.com/pokt-network/poktroll/testutil/integration"
	testutil "github.com/pokt-network/poktroll/testutil/integration"
	"github.com/pokt-network/poktroll/testutil/testrelayer"
	"github.com/pokt-network/poktroll/x/shared"
)

// TODO_UPNEXT(@Olshansk, #571): Implement these tests

func init() {
	cmd.InitSDKConfig()
}

func TestUpdateRelayMiningDifficulty_NewServiceSeenForTheFirstTime(t *testing.T) {
	var claimWindowOpenBlockHash, proofWindowOpenBlockHash, proofPathSeedBlockHash []byte

	// Create a new integration app
	integrationApp := integration.NewCompleteIntegrationApp(t)

	// Move forward a few blocks to move away from the genesis block
	integrationApp.NextBlocks(t, 3)

	// Get the current session and shared params
	session := getSession(t, integrationApp)
	sharedParams := getSharedParams(t, integrationApp)

	// Prepare the trie with a single mined relay
	trie := prepareSMST(t, integrationApp.GetSdkCtx(), integrationApp, session)

	// Compute the number of blocks to wait between different events
	// TODO_BLOCKER(@bryanchriswhite): See this comment: https://github.com/pokt-network/poktroll/pull/610#discussion_r1645777322
	sessionEndHeight := session.Header.SessionEndBlockHeight
	earliestSupplierClaimCommitHeight := shared.GetEarliestSupplierClaimCommitHeight(
		&sharedParams,
		sessionEndHeight,
		claimWindowOpenBlockHash,
		integrationApp.DefaultSupplier.GetAddress(),
	)
	earliestSupplierProofCommitHeight := shared.GetEarliestSupplierProofCommitHeight(
		&sharedParams,
		sessionEndHeight,
		proofWindowOpenBlockHash,
		integrationApp.DefaultSupplier.GetAddress(),
	)
	proofWindowCloseHeight := shared.GetProofWindowCloseHeight(&sharedParams, sessionEndHeight)

	// Wait until the earliest claim commit height.
	currentBlockHeight := integrationApp.GetSdkCtx().BlockHeight()
	numBlocksUntilClaimWindowOpenHeight := earliestSupplierClaimCommitHeight - currentBlockHeight
	require.Greater(t, numBlocksUntilClaimWindowOpenHeight, int64(0), "unexpected non-positive number of blocks until the earliest claim commit height")
	integrationApp.NextBlocks(t, int(numBlocksUntilClaimWindowOpenHeight))

	// Construct a new create claim message and commit it.
	createClaimMsg := proof.MsgCreateClaim{
		SupplierAddress: integrationApp.DefaultSupplier.Address,
		SessionHeader:   session.Header,
		RootHash:        trie.Root(),
	}
	result := integrationApp.RunMsg(t,
		&createClaimMsg,
		integration.WithAutomaticFinalizeBlock(),
		integration.WithAutomaticCommit(),
	)
	require.NotNil(t, result, "unexpected nil result when submitting a MsgCreateClaim tx")

	// Wait until the proof window is open
	currentBlockHeight = integrationApp.GetSdkCtx().BlockHeight()
	numBlocksUntilProofWindowOpenHeight := earliestSupplierProofCommitHeight - currentBlockHeight
	require.Greater(t, numBlocksUntilProofWindowOpenHeight, int64(0), "unexpected non-positive number of blocks until the earliest proof commit height")
	integrationApp.NextBlocks(t, int(numBlocksUntilProofWindowOpenHeight))

	// Construct a new proof message and commit it
	createProofMsg := proof.MsgSubmitProof{
		SupplierAddress: integrationApp.DefaultSupplier.Address,
		SessionHeader:   session.Header,
		Proof:           getProof(t, trie, proofPathSeedBlockHash, session.GetHeader().GetSessionId()),
	}
	result = integrationApp.RunMsg(t,
		&createProofMsg,
		integration.WithAutomaticFinalizeBlock(),
		integration.WithAutomaticCommit(),
	)
	require.NotNil(t, result, "unexpected nil result when submitting a MsgSubmitProof tx")

	// Wait until the proof window is closed
	currentBlockHeight = integrationApp.GetSdkCtx().BlockHeight()
	numBlocksUntilProofWindowCloseHeight := proofWindowCloseHeight - currentBlockHeight
	require.Greater(t, numBlocksUntilProofWindowOpenHeight, int64(0), "unexpected non-positive number of blocks until the earliest proof commit height")
	// TODO_TECHDEBT(@bryanchriswhite): Olshansky is unsure why the +1 is necessary here
	// but it was required to pass the test.
	integrationApp.NextBlocks(t, int(numBlocksUntilProofWindowCloseHeight)+1)

	// The number 14 was determined empirically by running the tests and will need
	// to be updated if they are changed.
	expectedNumEvents := 15
	// Check the number of events is consistent.
	events := integrationApp.GetSdkCtx().EventManager().Events()
	require.Equalf(t, expectedNumEvents, len(events), "unexpected number of total events")

	relayMiningEvents := testutilevents.FilterEvents[*tokenomics.EventRelayMiningDifficultyUpdated](t,
		events, "poktroll.tokenomics.EventRelayMiningDifficultyUpdated")
	require.Len(t, relayMiningEvents, 1, "unexpected number of relay mining difficulty updated events")
	relayMiningEvent := relayMiningEvents[0]
	require.Equal(t, "svc1", relayMiningEvent.ServiceId)
	// The default difficulty)
	require.Equal(t, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", relayMiningEvent.PrevTargetHashHexEncoded)
	require.Equal(t, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", relayMiningEvent.NewTargetHashHexEncoded)
	// The previous EMA is the same as the current one if the service is new
	require.Equal(t, uint64(1), relayMiningEvent.PrevNumRelaysEma)
	require.Equal(t, uint64(1), relayMiningEvent.NewNumRelaysEma)
}

func UpdateRelayMiningDifficulty_UpdatingMultipleServicesAtOnce(t *testing.T) {}

func UpdateRelayMiningDifficulty_UpdateServiceIsNotSeenForAWhile(t *testing.T) {}

func UpdateRelayMiningDifficulty_UpdateServiceIsIncreasing(t *testing.T) {}

func UpdateRelayMiningDifficulty_UpdateServiceIsDecreasing(t *testing.T) {}

// getSharedParams returns the shared parameters for the current block height.
func getSharedParams(t *testing.T, integrationApp *testutil.App) sharedtypes.Params {
	t.Helper()

	sharedQueryClient := sharedtypes.NewQueryClient(integrationApp.QueryHelper())
	sharedParamsReq := sharedtypes.QueryParamsRequest{}

	sharedQueryRes, err := sharedQueryClient.Params(integrationApp.GetSdkCtx(), &sharedParamsReq)
	require.NoError(t, err)

	return sharedQueryRes.Params
}

// getSession returns the current session for the default application and service.
func getSession(t *testing.T, integrationApp *testutil.App) *session.Session {
	t.Helper()

	sessionQueryClient := session.NewQueryClient(integrationApp.QueryHelper())
	getSessionReq := session.QueryGetSessionRequest{
		ApplicationAddress: integrationApp.DefaultApplication.Address,
		Service:            integrationApp.DefaultService,
		BlockHeight:        integrationApp.GetSdkCtx().BlockHeight(),
	}

	getSessionRes, err := sessionQueryClient.GetSession(integrationApp.GetSdkCtx(), &getSessionReq)
	require.NoError(t, err)
	require.NotNil(t, getSessionRes, "unexpected nil queryResponse")
	return getSessionRes.Session
}

// prepareSMST prepares an SMST with a single mined relay for the given session.
func prepareSMST(
	t *testing.T, ctx context.Context,
	integrationApp *testutil.App,
	session *session.Session,
) *smt.SMST {
	t.Helper()

	// Generating an ephemeral tree & spec just so we can submit
	// a proof of the right size.
	// TODO_TECHDEBT(#446): Centralize the configuration for the SMT spec.
	kvStore, err := badger.NewKVStore("")
	require.NoError(t, err)

	// NB: A signed mined relay is a MinedRelay type with the appropriate
	// payload, signatures and metadata populated.
	//
	// It does not (as of writing) adhere to the actual on-chain difficulty (i.e.
	// hash check) of the test service surrounding the scope of this test.
	minedRelay := testrelayer.NewSignedMinedRelay(t, ctx,
		session,
		integrationApp.DefaultApplication.Address,
		integrationApp.DefaultSupplier.Address,
		integrationApp.DefaultSupplierKeyringKeyringUid,
		integrationApp.GetKeyRing(),
		integrationApp.GetRingClient(),
	)

	trie := smt.NewSparseMerkleSumTrie(kvStore, sha256.New(), smt.WithValueHasher(nil))
	err = trie.Update(minedRelay.Hash, minedRelay.Bytes, 1)
	require.NoError(t, err)

	return trie
}

// getProof returns a proof for the given session for the empty path.
// If there is only one relay in the trie, the proof will be for that single
// relay since it is "closest" to any path provided, empty or not.
func getProof(
	t *testing.T,
	trie *smt.SMST,
	pathSeedBlockHash []byte,
	sessionId string,
) []byte {
	t.Helper()

	path := protocol.GetPathForProof(pathSeedBlockHash, sessionId)
	proof, err := trie.ProveClosest(path)
	require.NoError(t, err)

	proofBz, err := proof.Marshal()
	require.NoError(t, err)

	return proofBz
}
