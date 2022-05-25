package user

type Service interface {
	Get(id uint) (*Model, error)
	Delete(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Login(data LoginDTO) (*Model, error)
}

type service struct {
	repo Repository
}

var _ Service = service{}

func NewService(repo Repository) Service {
	return service{repo: repo}
}

func (s service) Get(id uint) (*Model, error) {
	return s.repo.Get(id)
}
func (s service) Delete(id uint) (*Model, error) {
	return s.repo.Delete(id)
}

func (s service) Create(model Model) (uint, error) {
	return s.repo.Create(model)
}

func (s service) Login(data LoginDTO) (*Model, error) {
	return s.repo.Login(data)
}
