package robotorders

type Service interface {
	Save(robotorders Robotorder) (Robotorder, error)
	FindAllUserid(userid int) ([]Robotorder, error)
	FindAllKodeinvoice(kodeinvoice string) ([]Robotorder, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(robotorders Robotorder) (Robotorder, error) {
	users, err := s.repository.Save(robotorders)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindAllUserid(userid int) ([]Robotorder, error) {
	users, err := s.repository.FindAllUserid(userid)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindAllKodeinvoice(kodeinvoice string) ([]Robotorder, error) {
	users, err := s.repository.FindAllKodeinvoice(kodeinvoice)
	if err != nil {
		return users, err
	}

	return users, nil
}
