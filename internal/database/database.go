package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sql.Open("libsql", dburl)
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
