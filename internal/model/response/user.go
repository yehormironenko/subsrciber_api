package response

type User struct {
	UserId    uint   `json:"user_id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
