package repository
import (
	"database/sql"
	"tg-bot/internal/models"
)
func (r *repository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (telegram_id, username, city) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at`
	return r.db.QueryRow(query, user.TelegramID, user.Username, user.City).
		Scan(&user.ID, &user.CreatedAt)
}
func (r *repository) GetUserByTelegramID(telegramID int64) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, telegram_id, username, city, created_at 
		FROM users 
		WHERE telegram_id = $1`
	err := r.db.QueryRow(query, telegramID).
		Scan(&user.ID, &user.TelegramID, &user.Username, &user.City, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
func (r *repository) UpdateUserCity(telegramID int64, city string) error {
	query := `UPDATE users SET city = $1 WHERE telegram_id = $2`
	_, err := r.db.Exec(query, city, telegramID)
	return err
}
