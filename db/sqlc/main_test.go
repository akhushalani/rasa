package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const (
	dbSource = "postgresql://root:secret@localhost:5432/rasa?sslmode=disable"
)

func TestMain(m *testing.M) {
	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("cannot parse db source:", err)
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
