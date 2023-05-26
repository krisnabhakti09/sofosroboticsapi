package robotproductpartdevices

type RobotProductDeviceFormatter struct {
	ID                   int    `json:"id"`
	Idrobotmaster        int    `json:"idrobotmaster"`
	Idrobotpartdevice    int    `json:"idrobotpartdevice"`
	Namerobotmaster      string `json:"namerobotmaster"`
	Namerobotpartdevice  string `json:"namerobotpartdevice"`
	Namecomponent        string `json:"namecomponent"`
	Descriptioncomponent string `json:"descriptioncomponent"`
	Image                string `json:"image"`
	Price                int    `json:"price"`
	Unit                 string `json:"unit"`
	Required             string `json:"required"`
}

func FormatRobotProductDevice(robotmasters Robotproductpartdevice) RobotProductDeviceFormatter {
	formatter := RobotProductDeviceFormatter{
		ID:                   robotmasters.ID,
		Idrobotmaster:        robotmasters.Idrobotmaster,
		Idrobotpartdevice:    robotmasters.Idrobotpartdevice,
		Namerobotmaster:      robotmasters.Namerobotmaster,
		Namerobotpartdevice:  robotmasters.Namerobotpartdevice,
		Namecomponent:        robotmasters.Namecomponent,
		Descriptioncomponent: robotmasters.Descriptioncomponent,
		Image:                robotmasters.Image,
		Price:                robotmasters.Price,
		Unit:                 robotmasters.Unit,
		Required:             robotmasters.Required,
	}

	return formatter
}

func FormatAll(teachers []Robotproductpartdevice) []RobotProductDeviceFormatter {
	teachersFormatter := []RobotProductDeviceFormatter{}

	for _, teac := range teachers {
		teacherFormatter := FormatRobotProductDevice(teac)
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}
