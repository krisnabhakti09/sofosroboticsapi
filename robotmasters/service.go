package robotmasters

type Service interface {
	Recomendation() ([]Robotmaster, error)
	FindAllRobotic() ([]Robotmaster, error)
	FindAllAutomation() ([]Robotmaster, error)
	Byid(id int) (Robotmaster, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Recomendation() ([]Robotmaster, error) {
	users, err := s.repository.Recomendation()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindAllRobotic() ([]Robotmaster, error) {
	users, err := s.repository.FindAllRobotic()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindAllAutomation() ([]Robotmaster, error) {
	users, err := s.repository.FindAllAutomation()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) Byid(id int) (Robotmaster, error) {
	users, err := s.repository.Byid(id)
	if err != nil {
		return users, err
	}

	return users, nil
}
