package caching

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	postgrestestcontainers "github.com/testcontainers/testcontainers-go/modules/postgres"
	redistestcontainers "github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	rc             *redis.Client
	pgxConn        *pgx.Conn
	postgresDBName = "test"
	postgresDBUser = "user"
	postgresDBPass = "password"
)

func TestCachingSmall(t *testing.T) {
	setup(t, false)
	test(t)
}

func TestCachingBig(t *testing.T) {
	setup(t, true)
	test(t)

}

func test(t *testing.T) {

	var (
		pgxRepo    = NewUsersRepoPostgresImpl(pgxConn)
		cachedRepo = NewUsersRepoCached(pgxRepo, rc, time.Second*10)
	)

	_, _ = pgxRepo.GetByUsername(context.Background(), "Bob22")
	_, _ = pgxRepo.GetByUsername(context.Background(), "Bob33")
	_, _ = pgxRepo.GetByUsername(context.Background(), "Bob44")

	t.Run("postgres only", func(t *testing.T) {
		start := time.Now()
		users, err := pgxRepo.GetByUsername(context.Background(), "Bob22")
		require.NoError(t, err)

		d := time.Since(start)
		fmt.Printf("Time it took to get %d users from Postgres: %v\n", len(users), d)

		assert.Greater(t, len(users), 0)
	})

	t.Run("cache miss", func(t *testing.T) {
		start := time.Now()
		users, err := cachedRepo.GetByUsername(context.Background(), "Bob33")
		require.NoError(t, err)
		d := time.Since(start)
		fmt.Printf("Time it took to get %d users from Redis and Postgres (miss): %v\n", len(users), d)

		assert.Greater(t, len(users), 0)
	})
	t.Run("cache hit", func(t *testing.T) {
		start := time.Now()
		users, err := cachedRepo.GetByUsername(context.Background(), "Bob44")
		require.NoError(t, err)
		d := time.Since(start)
		fmt.Printf("Time it took to get %d users from Redis and Postgres (miss): %v\n", len(users), d)

		start = time.Now()
		users, err = cachedRepo.GetByUsername(context.Background(), "Bob44")
		require.NoError(t, err)
		d = time.Since(start)
		fmt.Printf("Time it took to get %d users from Redis (hit): %v\n", len(users), d)

		assert.Greater(t, len(users), 0)
	})
}

func setup(t *testing.T, big bool) {
	setupRedis(t)
	setupPostgres(t, big)
}

func setupPostgres(t *testing.T, big bool) {
	ctx := context.Background()

	initScript := "init-user-db.sql"
	if big {
		initScript = "init-user-db-big.sql"
	}

	postgresContainer, err := postgrestestcontainers.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgrestestcontainers.WithInitScripts(initScript),
		postgrestestcontainers.WithDatabase(postgresDBName),
		postgrestestcontainers.WithUsername(postgresDBUser),
		postgrestestcontainers.WithPassword(postgresDBPass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2)),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Printf("failed to terminate container: %s\n", err)
		}
	})

	cs, err := postgresContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	fmt.Println(cs)

	conn, err := pgx.Connect(context.Background(), cs)
	require.NoError(t, err)

	require.NoError(t, conn.Ping(context.Background()))

	pgxConn = conn
}

func setupRedis(t *testing.T) {
	redisContainer, err := redistestcontainers.Run(context.Background(),
		"docker.io/redis:7",
		redistestcontainers.WithSnapshotting(10, 1),
		redistestcontainers.WithLogLevel(redistestcontainers.LogLevelVerbose),
	)
	require.NoError(t, err)

	// Clean up the container
	t.Cleanup(func() {
		if err := redisContainer.Terminate(context.Background()); err != nil {
			log.Printf("failed to terminate container: %s\n", err)
		}
	})

	cs, err := redisContainer.ConnectionString(context.Background())
	require.NoError(t, err)

	cs = strings.Replace(cs, "redis://", "", 1)

	rc = redis.NewClient(&redis.Options{
		Addr: cs,
	})

	require.NoError(t, rc.Ping(context.Background()).Err())
}
