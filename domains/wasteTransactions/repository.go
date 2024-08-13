package wasteTransactions

type Repository interface {
}

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}
