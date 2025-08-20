package models
import "time"
type User struct {
	ID         int       `db:"id" json:"id"`
	TelegramID int64     `db:"telegram_id" json:"telegram_id"`
	Username   string    `db:"username" json:"username"`
	City       string    `db:"city" json:"city"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
