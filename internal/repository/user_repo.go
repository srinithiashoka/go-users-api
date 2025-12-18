package repository

import (
	"database/sql"
	"go-users-api/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(u models.User) (int64, error) {
	res, err := r.DB.Exec(
		`INSERT INTO users (name, dob) VALUES (?, ?)`,
		u.Name, u.Dob,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	var u models.User
	err := r.DB.QueryRow(
		`SELECT id, name, dob FROM users WHERE id = ?`, id,
	).Scan(&u.ID, &u.Name, &u.Dob)

	return u, err
}

func (r *UserRepository) ListUsers() ([]models.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, dob FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Dob); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(u models.User) error {
	_, err := r.DB.Exec(
		`UPDATE users SET name=?, dob=? WHERE id=?`,
		u.Name, u.Dob, u.ID,
	)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id=?`, id)
	return err
}
