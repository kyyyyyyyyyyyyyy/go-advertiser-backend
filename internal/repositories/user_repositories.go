package repositories

import (
	"database/sql"
	"go-advertiser-backend/internal/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindAll() ([]models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			is_active,
			role,
			created_at,
			updated_at
		FROM users
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.IsActive,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) FindByID(
	id uuid.UUID,
) (*models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password,
			role,
			is_active,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(
	email string,
) (*models.User, error) {

	query := `
		SELECT
			id,
			name,
			email,
			password,
			is_active,
			role,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(
	name string,
	email string,
	password string,
	role string,
) (*models.User, error) {

	query := `
		INSERT INTO users (
			name,
			email,
			password,
			role
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			name,
			email,
			is_active,
			role,
			created_at,
			updated_at
	`

	var user models.User

	err := r.DB.QueryRow(
		query,
		name,
		email,
		password,
		role,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.IsActive,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
