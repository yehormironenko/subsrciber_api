package response

type Subscriber struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
