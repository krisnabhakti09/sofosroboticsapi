package robotorders

type Input struct {
	Idrobotproductpartdevice int    `json:"idrobotproductpartdevice" binding:"required"`
	Userid                   int    `json:"userid" binding:"required"`
	Kodeinvoice              string `json:"kodeinvoice" binding:"required"`
	Namecustomer             string `json:"namecustomer" binding:"required"`
	Namarobotmaster          string `json:"namerobotmaster" binding:"required"`
	Addresscustomer          string `json:"addresscustomer" binding:"required"`
	Delivery                 string `json:"delivery" binding:"required"`
	Pricedelivery            int    `json:"pricedelivery" binding:"required"`
	Pricemethod              string `json:"pricemethod" binding:"required"`
	Pricemethodsn            string `json:"pricemethodsn" binding:"required"`
	Subtotal                 int    `json:"subtotal" binding:"required"`
	Ppn                      int    `json:"ppn" binding:"required"`
	Totals                   int    `json:"totals" binding:"required"`
	Pesan                    string `json:"pesan"`
	Recommended              string `json:"recommended"`
	Deliverydesc             string `json:"deliverydesc" binding:"required"`
	Pricemethodadmin         int    `json:"pricemethodadmin" binding:"required"`
}

type Userid struct {
	Userid int `json:"userid" binding:"required"`
}

type Kodeinvoice struct {
	Kodeinvoice string `json:"kodeinvoice" binding:"required"`
}
