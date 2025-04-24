package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/bytedance/sonic"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/turahe/interpesona-data/internal/db/model"
	"github.com/turahe/interpesona-data/internal/helper/cache"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	AddUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	GetUserByPhone(ctx context.Context, phone string) (model.User, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	GetUsersWithPagination(ctx context.Context, limit int, offset int) ([]model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (bool, error)
	IsUserEmailExist(ctx context.Context, email string) (bool, error)
	IsUserPhoneExist(ctx context.Context, phone string) (bool, error)
}

type UserRepositoryImpl struct {
	pgxPool     *pgxpool.Pool
	redisClient redis.Cmdable
}

func NewUserRepository(pgxPool *pgxpool.Pool, redisClient redis.Cmdable) UserRepository {
	return &UserRepositoryImpl{
		pgxPool:     pgxPool,
		redisClient: redisClient,
	}
}

func (u *UserRepositoryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	key := "users"

	data, err := cache.Remember(ctx, key, 10*time.Minute, func() ([]byte, error) {
		var users []model.User
		rows, err := u.pgxPool.Query(ctx, "SELECT id, username, email, phone FROM users")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var user model.User
			err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}

		// Serialize users to bytes using Sonic
		userBytes, err := sonic.Marshal(users)
		if err != nil {
			return nil, err
		}

		return userBytes, nil
	})

	if err != nil {
		return nil, err
	}

	// Deserialize data to []model.User
	var users []model.User
	err = sonic.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User
	err := u.pgxPool.QueryRow(ctx, "SELECT id, username, email, phone FROM users WHERE id = $1 AND deleted_at IS NULL ", id).Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetUsersWithPagination(ctx context.Context, limit int, offset int) ([]model.User, error) {
	var users []model.User
	var totalUsers uint64
	fmt.Println(ctx.Value("limit"))

	rows, err := u.pgxPool.Query(ctx, "SELECT id, username, email, phone FROM users LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = u.pgxPool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&totalUsers)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Add user with transaction and return id
func (u *UserRepositoryImpl) AddUser(ctx context.Context, user model.User) (model.User, error) {
	tx, err := u.pgxPool.Begin(ctx)
	if err != nil {
		return model.User{}, err
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, "INSERT INTO users (id, username, email, phone) VALUES ($1, $2, $3, $4) RETURNING id, username, email, phone, created_at, updated_at", uuid.New(), user.UserName, user.Email, user.Phone).
		Scan(&user.ID, &user.UserName, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return model.User{}, err
	}

	// Delete cache
	err = cache.Remove(ctx, "users")
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	tx, err := u.pgxPool.Begin(ctx)
	if err != nil {
		return model.User{}, err
	}
	defer tx.Rollback(ctx)

	_, err = u.pgxPool.Exec(ctx, "UPDATE users SET username = $2, email = $3, phone = $4 WHERE id = $1", user.ID, user.UserName, user.Email, user.Phone)
	if err != nil {
		return model.User{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return model.User{}, err
	}

	// Delete cache
	err = cache.Remove(ctx, "users")
	if err != nil {
		return model.User{}, err
	}

	// Return the updated user
	return user, nil
}

func (u *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := u.pgxPool.QueryRow(ctx, "SELECT id, username, email, phone FROM users WHERE email = $1", email).Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	var user model.User
	err := u.pgxPool.QueryRow(ctx, "SELECT id, username, email, phone FROM users WHERE phone = $1", phone).Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	err := u.pgxPool.QueryRow(ctx, "SELECT id, username, email, phone FROM users WHERE username = $1", username).Scan(&user.ID, &user.UserName, &user.Email, &user.Phone)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) DeleteUser(ctx context.Context, id uuid.UUID) (bool, error) {
	tx, err := u.pgxPool.Begin(ctx)
	if err != nil {
		return false, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "UPDATE users SET deleted_at = NOW() WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return false, err
	}

	// Delete cache
	err = cache.Remove(ctx, "users")

	return true, nil

}

func (u *UserRepositoryImpl) IsUserEmailExist(ctx context.Context, email string) (bool, error) {
	var count int
	err := u.pgxPool.QueryRow(ctx, "SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
func (u *UserRepositoryImpl) IsUserPhoneExist(ctx context.Context, phone string) (bool, error) {
	var count int
	err := u.pgxPool.QueryRow(ctx, "SELECT COUNT(*) FROM users WHERE phone = $1", phone).Scan(&count)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
