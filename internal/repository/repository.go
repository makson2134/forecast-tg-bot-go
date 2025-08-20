package repository
import (
	"database/sql"
	"tg-bot/internal/models"
)
type Repository interface {
	CreateUser(user *models.User) error
	GetUserByTelegramID(telegramID int64) (*models.User, error)
	UpdateUserCity(telegramID int64, city string) error
}
type repository struct {
	db *sql.DB
}
func New(db *sql.DB) Repository {
	return &repository{db: db}
}
