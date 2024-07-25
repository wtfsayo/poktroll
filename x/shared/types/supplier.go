package types

// SupplierNotUnstaking is the value of `unstake_session_end_height` if the
// supplier is not actively in the unbonding period.
const SupplierNotUnstaking uint64 = 0

// IsUnbonding returns true if the supplier is actively unbonding.
// It determines if the supplier has submitted an unstake message, in which case
// the supplier has its UnstakeSessionEndHeight set.
func (s *Supplier) IsUnbonding() bool {
	return s.UnstakeSessionEndHeight != SupplierNotUnstaking
}

// IsActive returns whether the supplier is allowed to serve requests for the
// given serviceId and query height.
// A supplier is active for a given service starting from the session following
// the one in which the supplier staked for that service.
// A supplier that has submitted an unstake message is active until the end of
// the session containing the height at which unstake message was submitted.
func (s *Supplier) IsActive(queryHeight uint64, serviceId string) bool {
	if s.ServicesActivationHeight[serviceId] > queryHeight {
		return false
	}

	if s.IsUnbonding() {
		return queryHeight <= s.UnstakeSessionEndHeight
	}

	return true
}
