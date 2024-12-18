package ads

type Storage interface {
	CreateAd(title, description string, price float64) (int, error)
	GetAllAd() ([]Ad, error)
	UpdateAd(ad Ad) (bool, error)
	GetAd(id int) (*Ad, error)
	// TODO: add other CRUD methods
}

type Ad struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Service struct {
	store Storage
}

func NewService(store Storage) *Service {
	return &Service{store: store}
}

func (s *Service) CreateAd(title, description string, price float64) (*Ad, error) {
	id, err := s.store.CreateAd(title, description, price)
	if err != nil {
		return nil, err
	}
	return &Ad{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       price,
	}, nil
}

func (s *Service) GetAd(id int) (*Ad, error) {
	return s.store.GetAd(id)
}

func (s *Service) GetAllAds() ([]Ad, error) {
	ads, err := s.store.GetAllAd()
	if err != nil {
		return []Ad{}, err
	}
	return ads, nil
}

func (s *Service) UpdateAd(ad Ad) (bool, error) {
	return s.store.UpdateAd(ad)
}

// TODO: implement other CRUD methods
