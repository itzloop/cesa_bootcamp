package caching

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

type UsersRepo interface {
	GetByUsername(ctx context.Context, name string) ([]*User, error)
}

type UsersRepoPostgresImpl struct {
	conn *pgx.Conn
}

func NewUsersRepoPostgresImpl(conn *pgx.Conn) *UsersRepoPostgresImpl {
	return &UsersRepoPostgresImpl{
		conn: conn,
	}
}

func (repo *UsersRepoPostgresImpl) GetByUsername(ctx context.Context, name string) ([]*User, error) {
	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		users []*User
	)

	for rows.Next() {
		user := new(User)
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

type UsersRepoCached struct {
	mainRepo   UsersRepo
	cache      *redis.Client
	expiration time.Duration
}

func NewUsersRepoCached(mainRepo UsersRepo, cache *redis.Client, expiration time.Duration) *UsersRepoCached {
	return &UsersRepoCached{
		mainRepo:   mainRepo,
		cache:      cache,
		expiration: expiration,
	}
}

func (repo *UsersRepoCached) GetByUsername(ctx context.Context, name string) ([]*User, error) {
	// check if key exists in redis
	usersStr, err := repo.cache.GetEx(ctx, name, repo.expiration).Result()
	if err == nil {
		// cache hit
		var users = make([]*User, 10)
		if err = json.Unmarshal([]byte(usersStr), &users); err != nil {
			return nil, err
		}

		return users, nil
	}

	// something went wrong
	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	// cache miss
	users, err := repo.mainRepo.GetByUsername(ctx, name)
	if err != nil {
		return nil, err
	}

	// set to cache and return the results
	if err = repo.cacheUsers(ctx, name, users); err != nil {
		log.Printf("failed to cache users: %v\n", err)
	}

	return users, nil
}

func (repo *UsersRepoCached) cacheUsers(ctx context.Context, name string, users []*User) error {
	v, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return repo.cache.SetEx(ctx, name, string(v), repo.expiration).Err()
}
