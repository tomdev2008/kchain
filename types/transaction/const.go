package transaction

type TransactionType string

const (
	MetaDataAdd TransactionType = "addMetadata"
	MetaDataGet TransactionType = "getMetadata"
	AccountSet TransactionType = "setAccount"
	AccountGet TransactionType = "getAccount"
	DbSet string = "DbSet"
	DbGet string = "DbGet"
)
