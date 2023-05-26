package robotproductpartdevices

type GetInput struct {
	Idrobotmaster     int `json:"idrobotmaster" binding:"required"`
	Idrobotpartdevice int `json:"idrobotpartdevice" binding:"required"`
}

type GetID struct {
	ID int `json:"idrobotproductpartdevice" binding:"required"`
}
