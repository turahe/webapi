package repository

import (
	"github.com/turahe/interpesona-data/internal/db/pgx"
	"github.com/turahe/interpesona-data/internal/db/rdb"
)

type Repository struct {
	User UserRepository
	Job  JobRepository
}

func NewRepository() *Repository {
	pgxPool := pgx.GetPgxPool()
	redisClient := rdb.GetRedisClient()

	return &Repository{
		User: NewUserRepository(pgxPool, redisClient),
		Job:  NewJobRepository(pgxPool),
	}
}
