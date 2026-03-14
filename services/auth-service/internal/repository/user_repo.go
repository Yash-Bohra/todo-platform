package repository

import (
	"auth-service/internal/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := `
	INSERT INTO users(email,password_hash)
	VALUES($1,$2)
	RETURNING id
	`

	return r.DB.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	var user models.User

	query := `
	SELECT id,email,password_hash
	FROM users
	WHERE email=$1
	`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
