package request

type User struct {
	Email string `json:"email" validate:"required,email"`
}
