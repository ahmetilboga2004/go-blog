package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/ahmetilboga2004/go-blog/internal/models"
	"github.com/ahmetilboga2004/go-blog/pkg/utils"
	"github.com/google/uuid"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	userID := uuid.New()

	salt, err := utils.GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword := utils.HashPassword(user.Password, salt)

	query := `INSERT INTO users (id, firstName, lastName, username, email, password, salt) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, firstName, lastName, username, email, password, salt`
	err = r.DB.QueryRow(query, userID, user.FirstName, user.LastName, user.Username, user.Email, hashedPassword, salt).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Salt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetAll() ([]*models.User, error) {
	rows, err := r.DB.Query("SELECT id, firstName, lastName, username, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	query := `SELECT id, firstName, lastName, username, email FROM users WHERE id = ?`
	rows := r.DB.QueryRow(query, id)
	var user models.User
	if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsernameOrEmail(username, email string) (*models.User, error) {
	query := `SELECT id, firstName, lastName, username, email, password, salt FROM users WHERE username = ? OR email = ?`
	user := &models.User{}
	err := r.DB.QueryRow(query, username, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Salt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(id uuid.UUID, user *models.User) (*models.User, error) {
	query := `UPDATE users SET firstName = ?, lastName = ?, username = ?, email = ?, password = ? WHERE id = ? RETURNING id, firstName, lastName, username, email`
	row := r.DB.QueryRow(query, user.FirstName, user.LastName, user.Username, user.Email, user.Password, id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsEffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
