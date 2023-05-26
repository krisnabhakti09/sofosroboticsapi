package robotproductpartdevices

type Service interface {
	FindAllIDRobotProductPartDevice(robotproductpartdevice int) (Robotproductpartdevice, error)
	FindAllIDRobotMasterAndRobotPartDevice(robotmasters int, robotpartdevice int) ([]Robotproductpartdevice, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllIDRobotMasterAndRobotPartDevice(robotmasters int, robotpartdevice int) ([]Robotproductpartdevice, error) {
	users, err := s.repository.FindAllIDRobotMasterAndRobotPartDevice(robotmasters, robotpartdevice)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) FindAllIDRobotProductPartDevice(robotproductpartdevice int) (Robotproductpartdevice, error) {
	users, err := s.repository.FindAllIDRobotProductPartDevice(robotproductpartdevice)
	if err != nil {
		return users, err
	}

	return users, nil
}
