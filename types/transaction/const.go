package transaction

type TransactionType string

const (
	DbSet = "DbSet"
	DbGet = "DbGet"

	ValidatorSet = "ValidatorSet"

	AccountSet = "AccountSet"
)
