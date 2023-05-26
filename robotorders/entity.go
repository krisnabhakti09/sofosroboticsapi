package robotorders

type Robotorder struct {
	ID                       int
	Idrobotproductpartdevice int
	Userid                   int
	Kodeinvoice              string
	Namecustomer             string
	Namarobotmaster          string
	Addresscustomer          string
	Delivery                 string
	Pricedelivery            int
	Pricemethod              string
	Pricemethodsn            string
	Subtotal                 int
	Ppn                      int
	Totals                   int
	Pesan                    string
	Deliverydesc             string
	Pricemethodadmin         int
	Recommended              string
}
