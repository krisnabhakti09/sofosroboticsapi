package robotmasters

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Typemaster  string `json:"typemaster" binding:"required"`
	Price       string `json:"price" binding:"required"`
}

type IDINPUT struct {
	ID int `json:"id" binding:"required"`
}
