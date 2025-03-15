package constants

import "time"

// Application constants
const (
	// App information
	AppName    = "MLM System"
	AppVersion = "0.1.0"

	// Default timeout durations
	DefaultTimeout = 30 * time.Second
	DBTimeout      = 10 * time.Second

	// Maximum request body size (10MB)
	MaxBodySize = 10 * 1024 * 1024

	// Pagination defaults
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// User roles
const (
	RoleAdmin    = "admin"
	RoleManager  = "manager"
	RoleUpline   = "upline"
	RoleDownline = "downline"
	RoleFinance  = "finance"
)

// Order statuses
const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusShipped   = "shipped"
	OrderStatusDelivered = "delivered"
	OrderStatusCancelled = "cancelled"
)

// Commission statuses
const (
	CommissionStatusPending  = "pending"
	CommissionStatusApproved = "approved"
	CommissionStatusPaid     = "paid"
	CommissionStatusRejected = "rejected"
)

// Commission types
const (
	CommissionTypeDirect  = "direct"
	CommissionTypeTeam    = "team"
	CommissionTypeBonus   = "bonus"
	CommissionTypeOveride = "override"
)

// Transaction types
const (
	TransactionTypeCommission = "commission"
	TransactionTypeWithdrawal = "withdrawal"
	TransactionTypeBonus      = "bonus"
	TransactionTypeRefund     = "refund"
)

// Error messages
const (
	ErrInvalidCredentials    = "Invalid username or password"
	ErrUserNotFound          = "User not found"
	ErrProductNotFound       = "Product not found"
	ErrOrderNotFound         = "Order not found"
	ErrUnauthorized          = "Unauthorized access"
	ErrInsufficientInventory = "Insufficient inventory"
	ErrInvalidInput          = "Invalid input data"
	ErrInternalServer        = "Internal server error"
	ErrDuplicateEntry        = "Duplicate entry"
)