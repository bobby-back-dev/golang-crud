package repository

import (
	"context"
	"fmt"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: conn,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {

	query := `INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email`
	ctx := context.Background()

	users := &models.User{}
	if err := ur.pool.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&users.ID, &users.Name, &users.Email); err != nil {
		fmt.Println("gagal membuat data user")
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetAll(user *models.User) (*models.User, error) {

	query := `SELECT id, name, email, password FROM users`
	ctx := context.Background()
	users := &models.User{}

	data, err := ur.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		err := data.Scan(&users.ID, &users.Name, &users.Email, &users.Password)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {

	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	ctx := context.Background()
	users := &models.User{}

	if err := ur.pool.QueryRow(ctx, query, id).Scan(&users.ID, &users.Name, &users.Email); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) UpdateUser(user *models.User) (*models.User, error) {

	query := `UPDATE users SET name = $1, email = $2, password = $3`
	ctx := context.Background()
	users := &models.User{}

	if err := ur.pool.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&users.Name, &users.Email); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) error {

	Query := `DELETE FROM users WHERE id = $1`
	ctx := context.Background()
	_, err := ur.pool.Exec(ctx, Query, id)
	if err != nil {
		return err
	}
	return nil
}
