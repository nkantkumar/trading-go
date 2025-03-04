package order

type Repository interface {
	Save(order *Order) error
	GetAll() ([]Order, error)
}
