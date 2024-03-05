package repository

import (
	"auth-go/services/user/internal/domain"
	"database/sql"
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
	stmt := "INSERT INTO users(email, password, created_at) VALUES(?, ?, ?);SELECT LAST_INSERTED_ID();"
	err := r.DB.QueryRow(stmt, userToInsert.Email, userToInsert.Password, userToInsert.CreatedAt).Scan(&id)
	if err != nil {
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

func (r *UserRepositoryImpl) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	stmt := "SELECT email, password FROM users WHERE email=?"
	if err := r.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword); err != nil {
		return -1, err
	}

	return id, nil
}
