package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type UserRepo interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, username string) (*models.User, error)
}
type UserRepository struct {
	pool *pgxpool.Pool
	hash *crypto.Hash
}

func NewUserRepository(conn *pgxpool.Pool, hash *crypto.Hash) *UserRepository {
	return &UserRepository{
		pool: conn,
		hash: hash,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {

	query := `INSERT INTO users(username, password_hash, display_name) VALUES ($1, $2, $3) RETURNING id, created_at`

	if err := ur.pool.QueryRow(ctx, query, user.Username, user.PasswordHash, user.DisplayName).Scan(&user.ID, &user.CreatedAt); err != nil {
		fmt.Println("gagal membuat data user")
		return nil, err
	}
	return user, nil
}
func (ur *UserRepository) Login(ctx context.Context, username string) (*models.User, error) {

	log.Printf("Attempting to find user with email: %s", username)
	query := `SELECT id, username, display_name, password_hash,  created_at FROM users WHERE username = $1`
	users := &models.User{}

	if err := ur.pool.QueryRow(ctx, query, username).Scan(&users.ID, &users.Username, &users.DisplayName, &users.PasswordHash, &users.CreatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return users, nil
}

//func (ur *UserRepository) GetAll() (*[]models.User, error) {
//	query := `SELECT id, name, email FROM users`
//	ctx := context.Background()
//
//	rows, err := ur.pool.Query(ctx, query)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var allUsers []models.User
//
//	for rows.Next() {
//		var user models.User
//		err := rows.Scan(&user.ID, &user.Name, &user.Email)
//		if err != nil {
//			return nil, err
//		}
//
//		allUsers = append(allUsers, user)
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//	return &allUsers, nil
//}
//
//func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
//
//	query := `SELECT id, name, email FROM users WHERE id = $1`
//	ctx := context.Background()
//	users := &models.User{}
//
//	if err := ur.pool.QueryRow(ctx, query, id).Scan(&users.ID, &users.Name, &users.Email); err != nil {
//		return nil, err
//	}
//	return users, nil
//}
//
//func (ur *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
//
//	query := `UPDATE users SET name = $1, email = $2, password = $3`
//	ctx := context.Background()
//	users := &models.User{}
//
//	if err := ur.pool.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&users.Name, &users.Email); err != nil {
//		return nil, err
//	}
//	return users, nil
//}
//
//func (ur *UserRepository) DeleteUser(id int) error {
//
//	Query := `DELETE FROM users WHERE id = $1`
//	ctx := context.Background()
//	_, err := ur.pool.Exec(ctx, Query, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
