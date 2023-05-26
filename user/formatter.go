package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"` // user dan admin
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Number:   user.Number,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Avatar:   user.Avatar,
		Token:    token,
	}

	return formatter
}

func FormatAll(teachers []User) []UserFormatter {
	teachersFormatter := []UserFormatter{}

	for _, teac := range teachers {
		teacherFormatter := FormatUser(teac, "")
		teachersFormatter = append(teachersFormatter, teacherFormatter)

	}

	return teachersFormatter
}
