package keeper

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/pokt-network/poktroll/proto/types/proof"
	"github.com/pokt-network/poktroll/proto/types/session"
	sharedtypes "github.com/pokt-network/poktroll/proto/types/shared"
	"github.com/pokt-network/poktroll/x/shared"
)

type msgWithSessionAndSupplier interface {
	GetSessionHeader() *session.SessionHeader
	GetSupplierAddress() string
}

// queryAndValidateSessionHeader ensures that a session with the sessionID of the given session
// header exists and that this session includes the supplier with the given address.
// It returns a session which is hydrated with the on-chain session data.
func (k msgServer) queryAndValidateSessionHeader(
	ctx context.Context,
	msg msgWithSessionAndSupplier,
) (*session.Session, error) {
	logger := k.Logger().With("method", "queryAndValidateSessionHeader")

	sessionHeader := msg.GetSessionHeader()
	supplierAddr := msg.GetSupplierAddress()

	sessionReq := &session.QueryGetSessionRequest{
		ApplicationAddress: sessionHeader.GetApplicationAddress(),
		Service:            sessionHeader.GetService(),
		BlockHeight:        sessionHeader.GetSessionStartBlockHeight(),
	}

	// Get the on-chain session for the ground-truth against which the given
	// session header is to be validated.
	sessionRes, err := k.Keeper.sessionKeeper.GetSession(ctx, sessionReq)
	if err != nil {
		return nil, err
	}
	onChainSession := sessionRes.GetSession()

	logger.
		With(
			"session_id", onChainSession.GetSessionId(),
			"session_end_height", sessionHeader.GetSessionEndBlockHeight(),
			"supplier", supplierAddr,
		).
		Debug("got sessionId for proof")

	// Ensure that the given session header's session ID matches the on-chain onChainSession ID.
	if sessionHeader.GetSessionId() != onChainSession.GetSessionId() {
		return nil, proof.ErrProofInvalidSessionId.Wrapf(
			"session ID does not match on-chain session ID; expected %q, got %q",
			onChainSession.GetSessionId(),
			sessionHeader.GetSessionId(),
		)
	}

	// NB: it is redundant to assert that the service ID in the request matches the
	// on-chain session service ID because the session is queried using the service
	// ID as a parameter. Either a different session (i.e. different session ID)
	// or an error would be returned depending on whether an application/supplier
	// pair exists for the given service ID or not, respectively.

	// Ensure the given supplier is in the onChainSession supplier list.
	if isSupplerFound := foundSupplier(
		sessionRes.GetSession().GetSuppliers(),
		supplierAddr,
	); !isSupplerFound {
		return nil, proof.ErrProofNotFound.Wrapf(
			"supplier address %q not found in session ID %q",
			supplierAddr,
			sessionHeader.GetSessionId(),
		)
	}

	return onChainSession, nil
}

// validateClaimWindow returns an error if the given session is not eligible for claiming.
// It *assumes* that the msg's session header is a valid on-chain session with correct
// height fields. First call #queryAndValidateSessionHeader to ensure any user-provided
// session header is valid and correctly hydrated.
func (k msgServer) validateClaimWindow(
	ctx context.Context,
	msg *proof.MsgCreateClaim,
) error {
	logger := k.Logger().With("method", "validateClaimWindow")
	sessionHeader := msg.GetSessionHeader()
	sharedParams := k.sharedKeeper.GetParams(ctx)

	sessionEndHeight := sessionHeader.GetSessionEndBlockHeight()

	// Get the claim window open and close heights for the given session header.
	claimWindowOpenHeight := shared.GetClaimWindowOpenHeight(&sharedParams, sessionEndHeight)
	claimWindowCloseHeight := shared.GetClaimWindowCloseHeight(&sharedParams, sessionEndHeight)

	// Get the earliest claim commit height for the given supplier.
	earliestClaimCommitHeight, err := k.sharedQuerier.GetEarliestSupplierClaimCommitHeight(
		ctx,
		sessionEndHeight,
		msg.GetSupplierAddress(),
	)
	if err != nil {
		return err
	}

	// Get the current block height.
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()

	// Ensure the current block height is ON or AFTER the supplier's earliest claim commit height.
	// TODO_BLOCKER(@bryanchriswhite, @red-0ne): Enforce an additional "latest
	// supplier claim/proof commit offset" such that all suppliers have the same
	// "supplier claim/proof commit window" size.
	// See: https://github.com/pokt-network/poktroll/pull/620/files#r1656548473.
	if currentHeight < earliestClaimCommitHeight {
		return proof.ErrProofClaimOutsideOfWindow.Wrapf(
			"current block height (%d) is less than the session's earliest claim commit height (%d)",
			currentHeight,
			earliestClaimCommitHeight,
		)
	}

	// Ensure the current block height is BEFORE the claim window close height.
	if currentHeight > claimWindowCloseHeight {
		return proof.ErrProofClaimOutsideOfWindow.Wrapf(
			"current block height (%d) is greater than session claim window close height (%d)",
			currentHeight,
			claimWindowCloseHeight,
		)
	}

	logger.
		With(
			"current_height", currentHeight,
			"session_end_height", sessionEndHeight,
			"claim_window_open_height", claimWindowOpenHeight,
			"earliest_claim_commit_height", earliestClaimCommitHeight,
			"claim_window_close_height", claimWindowCloseHeight,
			"supplier_addr", msg.GetSupplierAddress(),
		).
		Debug("validated claim window")

	return nil
}

// validateProofWindow returns an error if the given session is not eligible for proving.
// It *assumes* that the msg's session header is a valid on-chain session with correct
// height fields. First call #queryAndValidateSessionHeader to ensure any user-provided
// session header is valid and correctly hydrated.
func (k msgServer) validateProofWindow(
	ctx context.Context,
	msg *proof.MsgSubmitProof,
) error {
	logger := k.Logger().With("method", "validateProofWindow")
	sessionHeader := msg.GetSessionHeader()
	sharedParams := k.sharedKeeper.GetParams(ctx)

	sessionEndHeight := sessionHeader.GetSessionEndBlockHeight()

	// Get the proof window open and close heights for the given session header.
	proofWindowOpenHeight := shared.GetProofWindowOpenHeight(&sharedParams, sessionEndHeight)
	proofWindowCloseHeight := shared.GetProofWindowCloseHeight(&sharedParams, sessionEndHeight)

	// Get the earliest proof commit height for the given supplier.
	earliestProofCommitHeight, err := k.sharedQuerier.GetEarliestSupplierProofCommitHeight(
		ctx,
		sessionEndHeight,
		msg.GetSupplierAddress(),
	)
	if err != nil {
		return err
	}

	// Get the current block height.
	sdkCtx := cosmostypes.UnwrapSDKContext(ctx)
	currentHeight := sdkCtx.BlockHeight()

	// Ensure the current block height is ON or AFTER the earliest proof commit height.
	if currentHeight < earliestProofCommitHeight {
		return proof.ErrProofProofOutsideOfWindow.Wrapf(
			"current block height (%d) is less than session's earliest proof commit height (%d)",
			currentHeight,
			earliestProofCommitHeight,
		)
	}

	// Ensure the current block height is BEFORE the proof window close height.
	if currentHeight > proofWindowCloseHeight {
		return proof.ErrProofProofOutsideOfWindow.Wrapf(
			"current block height (%d) is greater than session proof window close height (%d)",
			currentHeight,
			proofWindowCloseHeight,
		)
	}

	logger.
		With(
			"current_height", currentHeight,
			"session_end_height", sessionEndHeight,
			"proof_window_open_height", proofWindowOpenHeight,
			"earliest_proof_commit_height", earliestProofCommitHeight,
			"proof_window_close_height", proofWindowCloseHeight,
			"supplier_addr", msg.GetSupplierAddress(),
		).
		Debug("validated proof window")

	return nil
}

// foundSupplier ensures that the given supplier address is in the given list of suppliers.
func foundSupplier(suppliers []*sharedtypes.Supplier, supplierAddr string) bool {
	for _, supplier := range suppliers {
		if supplier.Address == supplierAddr {
			return true
		}
	}
	return false
}
