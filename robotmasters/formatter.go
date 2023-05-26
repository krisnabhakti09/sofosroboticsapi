package robotmasters

type RobotMasterFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Typemaster  string `json:"typemaster"`
	Price       int    `json:"price"`
}

func FormatRobotMaster(robotmasters Robotmaster, token string) RobotMasterFormatter {
	formatter := RobotMasterFormatter{
		ID:          robotmasters.ID,
		Name:        robotmasters.Name,
		Description: robotmasters.Description,
		Image:       robotmasters.Image,
		Typemaster:  robotmasters.Typemaster,
		Price:       robotmasters.Price,
	}

	return formatter
}

func FormatAll(teachers []Robotmaster) []RobotMasterFormatter {
	teachersFormatter := []RobotMasterFormatter{}

	for _, teac := range teachers {
		teacherFormatter := FormatRobotMaster(teac, "")
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}
