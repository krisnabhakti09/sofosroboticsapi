package robotorders

type RobotorderFormatter struct {
	ID                       int    `json:"id"`
	Idrobotproductpartdevice int    `json:"idrobotproductpartdevice"`
	Userid                   int    `json:"userid"`
	Kodeinvoice              string `json:"kodeinvoice"`
	Namecustomer             string `json:"namecustomer"`
	Namarobotmaster          string `json:"namerobotmaster"`
	Addresscustomer          string `json:"addresscustomer"`
	Delivery                 string `json:"delivery"`
	Pricedelivery            int    `json:"pricedelivery"`
	Pricemethod              string `json:"pricemethod"`
	Pricemethodsn            string `json:"pricemethodsn"`
	Subtotal                 int    `json:"subtotal"`
	Ppn                      int    `json:"ppn"`
	Totals                   int    `json:"totals"`
	Pesan                    string `json:"pesan"`
	Deliverydesc             string `json:"deliverydesc"`
	Pricemethodadmin         int    `json:"pricemethodadmin"`
	Recommended              string `json:"recommended"`
}

func FormatRobotorder(robotorders Robotorder) RobotorderFormatter {
	formatter := RobotorderFormatter{
		ID:                       robotorders.ID,
		Idrobotproductpartdevice: robotorders.Idrobotproductpartdevice,
		Userid:                   robotorders.Userid,
		Kodeinvoice:              robotorders.Kodeinvoice,
		Namecustomer:             robotorders.Namecustomer,
		Namarobotmaster:          robotorders.Namarobotmaster,
		Addresscustomer:          robotorders.Addresscustomer,
		Delivery:                 robotorders.Delivery,
		Pricedelivery:            robotorders.Pricedelivery,
		Pricemethod:              robotorders.Pricemethod,
		Pricemethodsn:            robotorders.Pricemethodsn,
		Subtotal:                 robotorders.Subtotal,
		Ppn:                      robotorders.Ppn,
		Totals:                   robotorders.Totals,
		Pesan:                    robotorders.Pesan,
		Deliverydesc:             robotorders.Deliverydesc,
		Pricemethodadmin:         robotorders.Pricemethodadmin,
		Recommended:              robotorders.Recommended,
	}

	return formatter
}

func FormatAll(teachers []Robotorder) []RobotorderFormatter {
	teachersFormatter := []RobotorderFormatter{}

	for _, teac := range teachers {
		teacherFormatter := FormatRobotorder(teac)
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}
