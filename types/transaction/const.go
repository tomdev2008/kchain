package transaction

type TransactionType string

const (
	DbSet = "DbSet"
	DbGet = "DbGet"

	ValidatorSet = "ValidatorSet"
	ValidatorDel = "ValidatorDel"

	AccountSet = "AccountSet"
	AccountGet = "AccountGet"
	AccountDel = "AccountDel"
)
