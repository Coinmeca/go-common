package repository

import (
	"context"
	"fmt"
	"net/url"
	"os"

	cv "github.com/coinmeca/go-common/chain"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	SCHEMA string
	CHAIN  int
	POOL   *pgxpool.Pool
)

func dbURL() string {
	//c := config.GetConfig()
	//fmt.Println("level is", c.LogLevel)

	if os.Getenv("POSTGRES_USER") == "" {
		godotenv.Load(".env")
	}

	fmt.Print(os.Getenv("POSTGRES_PORT"))

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres", os.Getenv("POSTGRES_USER"), url.QueryEscape(os.Getenv("POSTGRES_PASSWORD")), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))
	//url := "postgres://exa:exameca@34.64.62.105:5432/postgres"
	// url := "postgres://exa:exameca@localhost:5432/postgres"

	return url
}

func InitDB(ctx context.Context, schema string) {
	SetDBSchema(schema)

	uri := dbURL()
	conn, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}
	err = conn.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("unable to ping to database: %v", err))
	}
	POOL = conn
}

func CloseDB() {
	if POOL != nil {
		POOL.Close()
	}
}

func NewConnection(ctx context.Context, schema string) (*pgxpool.Pool, error) {
	SetDBSchema(schema)

	uri := dbURL()
	//logger.Info.Printf("Connecting Pool ...%v", uri)

	conn, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func SetDBSchema(schema string) {
	SCHEMA = schema
	CHAIN = cv.ChainNameMap[SCHEMA]
}
