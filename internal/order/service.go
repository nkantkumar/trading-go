package order

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) PlaceOrder(order *Order) error {
	order.Timestamp = time.Now()
	return s.repo.Save(order)
}

func (s *Service) ListOrders() ([]Order, error) {
	return s.repo.GetAll()
}
