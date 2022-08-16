package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}

	return formatter
}

func FormatUserRegister(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.PasswordHash,
		Token:    token,
	}

	return formatter
}
