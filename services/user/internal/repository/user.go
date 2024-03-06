package repository

import (
	"auth-go/services/user/internal/domain"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Insert(userToInsert domain.User) (int, error) {
	var id int
	stmt := "INSERT INTO users(email, password, created_at) VALUES(?, ?, ?);"
	_, err := r.DB.Exec(stmt, userToInsert.Email, userToInsert.Password, userToInsert.CreatedAt)
	if err != nil {
		return -1, err
	}
	if err := r.DB.QueryRow("SELECT LAST_INSERT_ID();").Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	stmt := "SELECT * FROM users"
	rows, err := r.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepositoryImpl) Get(id int) (domain.User, error) {
	var user domain.User
	stmt := "SELECT * FROM users WHERE id=?"
	err := r.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return nil
}

func (r *UserRepositoryImpl) Authenticate(email string, password []byte) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users WHERE email=?"
	if err := r.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword); err != nil {
		return -1, ErrEmailNotFound
	}
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return -1, ErrIncorrectPassword
	}
	return id, nil
}

func (r *UserRepositoryImpl) ExistsByEmail(email string) bool {
	var exists bool
	stmt := "SELECT EXISTS(SELECT TRUE FROM users WHERE email=?)"
	if err := r.DB.QueryRow(stmt, email).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (r *UserRepositoryImpl) Exists(id int) bool {
	var exists bool
	stmt := "SELECT EXISTS(SELECT TRUE FROM users WHERE id=?)"
	if err := r.DB.QueryRow(stmt, id).Scan(&exists); err != nil {
		return false
	}
	return exists
}
