package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Fullname:   user.Fullname,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
