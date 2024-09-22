package user

type UserFilter struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Status   Status `json:"status"`
}
