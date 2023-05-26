package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Number   string `json:"number" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	ID     int    `json:"userid" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required"`
	Email  string `json:"email" binding:"required"`
}
