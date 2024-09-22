package response

import "time"

type HealthCheck struct {
	Status         string    `json:"status"`
	DBStatus       string    `json:"db_status"`
	DBCheckedAt    time.Time `json:"db_checked_at"`
	RedisStatus    string    `json:"redis_status"`
	RedisCheckedAt time.Time `json:"redis_checked_at"`
}
