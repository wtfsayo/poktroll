package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ClaimList: []Claim{},
		ProofList: []Proof{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in claim
	claimPrimaryKeyMap := make(map[string]struct{})

	// Ensure claims are unique with respect to a given session ID and supplier address.
	for _, claim := range gs.ClaimList {
		// TODO_BLOCKER: ensure the corresponding supplier exists and is staked.

		if claim.GetRootHash() == nil {
			return fmt.Errorf("root hash cannot be nil")
		}

		if len(claim.GetRootHash()) == 0 {
			return fmt.Errorf("root hash cannot be empty")
		}

		sessionId := claim.GetSessionHeader().GetSessionId()
		primaryKey := string(ClaimPrimaryKey(sessionId, claim.SupplierAddress))
		if _, ok := claimPrimaryKeyMap[primaryKey]; ok {
			return fmt.Errorf("duplicated supplierAddr for claim")
		}
		claimPrimaryKeyMap[primaryKey] = struct{}{}
	}
	// Check for duplicated index in proof
	proofPrimaryKeyMap := make(map[string]struct{})

	for _, proof := range gs.ProofList {
		primaryKey := string(ProofPrimaryKey(
			proof.GetSessionHeader().GetSessionId(),
			proof.GetSupplierAddress(),
		))
		if _, ok := proofPrimaryKeyMap[primaryKey]; ok {
			return fmt.Errorf("duplicated primaryKey for proof")
		}
		proofPrimaryKeyMap[primaryKey] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
