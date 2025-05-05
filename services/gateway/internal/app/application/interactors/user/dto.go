package user

type GetUserByIdResponseDto struct {
	UserId       string `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	IsActive     bool   `json:"is_active"`
	LastActivity string `json:"last_activity"`
}
