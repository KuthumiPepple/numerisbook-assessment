package util

// Constants for all valid invoice status
const (
	DRAFT           = "draft"
	PENDING_PAYMENT = "pending_payment"
	OVERDUE         = "overdue"
	PAID            = "paid"
)

// IsValidStatus returns true if s is a valid invoice status
func IsValidStatus(s string) bool {
	switch s {
	case DRAFT, PENDING_PAYMENT, OVERDUE, PAID:
		return true
	}
	return false
}
