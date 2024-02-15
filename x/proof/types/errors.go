package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/proof module sentinel errors
var (
	ErrProofInvalidAddress = sdkerrors.Register(ModuleName, 2, "invalid address")
	//ErrSupplierUnauthorized = sdkerrors.Register(ModuleName, 3, "unauthorized supplier signer")
	//ErrSupplierNotFound                  = sdkerrors.Register(ModuleName, 4, "supplier not found")
	//ErrSupplierInvalidServiceConfig      = sdkerrors.Register(ModuleName, 5, "invalid service config")
	//ErrSupplierInvalidSessionStartHeight = sdkerrors.Register(ModuleName, 6, "invalid session start height")
	ErrProofInvalidSessionId = sdkerrors.Register(ModuleName, 7, "invalid session ID")
	//ErrSupplierInvalidService            = sdkerrors.Register(ModuleName, 8, "invalid service in supplier")
	//ErrSupplierInvalidClaimRootHash      = sdkerrors.Register(ModuleName, 9, "invalid root hash")
	ErrProofInvalidSessionEndHeight = sdkerrors.Register(ModuleName, 10, "invalid session ending height")
	ErrProofInvalidQueryRequest     = sdkerrors.Register(ModuleName, 11, "invalid query request")
	ErrProofClaimNotFound           = sdkerrors.Register(ModuleName, 12, "claim not found")
	//ErrSupplierProofNotFound             = sdkerrors.Register(ModuleName, 13, "proof not found")
	//ErrSupplierInvalidProof              = sdkerrors.Register(ModuleName, 14, "invalid proof")
	//ErrSupplierInvalidClosestMerkleProof = sdkerrors.Register(ModuleName, 15, "invalid closest merkle proof")
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
)
