package repository

import (
	"context"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {

	tesDbDsn := "postgres://bobby:bobby@localhost:5435/crud?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), tesDbDsn)
	require.NoError(t, err, "failed to create database connection")

	t.Cleanup(func() {
		_, err := pool.Exec(context.Background(), "TRUNCATE TABLE users RESTART IDENTITY CASCADE")
		if err != nil {
			log.Fatal("failed to truncate users table")
		}
		pool.Close()
	})
	return pool
}

func TestUserRepository_Create(t *testing.T) {
	// Setup
	pool := setupTestDB(t)
	hash := crypto.Hash{}
	userRepo := NewUserRepository(pool, &hash)
	ctx := context.Background()

	userToCreate := &models.User{
		Username:     "bobby",
		PasswordHash: "qwerty123",
		DisplayName:  "Bobby",
	}

	hashedPassword, err := hash.HashPassword(userToCreate.PasswordHash)
	require.NoError(t, err, "gagal hash password")
	userToCreate.PasswordHash = hashedPassword

	createdUser, err := userRepo.Create(ctx, userToCreate)

	assert.NoError(t, err, "Seharusnya tidak ada error saat membuat user")
	require.NotNil(t, createdUser, "User yang dikembalikan tidak boleh nil")

	assert.NotZero(t, createdUser.ID, "user id seharusnya tidak nol")
	assert.NotZero(t, createdUser.CreatedAt, "created at seharusnya tidak nol")

	var userFromDB models.User

	dbErr := pool.QueryRow(ctx,
		"SELECT id, username, display_name FROM users WHERE id=$1",
		createdUser.ID).Scan(
		&userFromDB.ID,
		&userFromDB.Username,
		&userFromDB.DisplayName,
	)

	assert.NoError(t, dbErr, "Gagal mengambil user yang baru dibuat dari DB")
	assert.Equal(t, createdUser.Username, userFromDB.Username)
	assert.Equal(t, createdUser.DisplayName, userFromDB.DisplayName)
}

func TestLogin(t *testing.T) {
	pool := setupTestDB(t)
	hash := crypto.Hash{}
	userRepo := NewUserRepository(pool, &hash)
	ctx := context.Background()
	passwordMentah := "qwerty123"
	hashedPassword, err := hash.HashPassword(passwordMentah)
	require.NoError(t, err)

	userToCreate := &models.User{
		Username:     "bobay",
		PasswordHash: hashedPassword,
		DisplayName:  "Bobby Display Name",
	}
	// Panggil method Create untuk memasukkan data ke DB tes.
	data, err := userRepo.Create(ctx, userToCreate)
	require.NoError(t, err, "Gagal membuat user untuk data tes")
	assert.NotNil(t, data, "User data tes")

	// 3. Act: Panggil method yang ingin kita tes.
	foundUser, err := userRepo.Login(ctx, "bobay")

	// 4. Assert: Periksa hasilnya.
	assert.NoError(t, err, "Seharusnya tidak ada error saat mencari user yang ada")
	require.NotNil(t, foundUser, "User yang ditemukan tidak boleh nil")

	assert.Equal(t, userToCreate.ID, foundUser.ID)
	assert.Equal(t, "bobay", foundUser.Username)
	assert.Equal(t, "Bobby Display Name", foundUser.DisplayName)
	assert.NotEmpty(t, foundUser.PasswordHash, "Password hash seharusnya tidak kosong")
}
