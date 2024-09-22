package response

import "time"

type UserRegister struct {
	UserID    string    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	AccessToken      string    `json:"access_token"`
	AccessExpiredAt  time.Time `json:"access_expired_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiredAt time.Time `json:"refresh_expired_at"`
}
