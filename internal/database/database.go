package database

import (
	"context"
	"database/sql"
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

// Write functions related to DB like checking health, querying data, here

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		return map[string]string{
			"error": "DB is Down",
		}
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

/* You can refer to this demo function :

func (s *service) GetUserByID(id string) string {
	stmt := `SELECT username FROM users WHERE id= ? `
	row := s.db.QueryRow(stmt, id)
	var userName string
	err := row.Scan(&userName)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows selected")
			return ""
		}
		log.Fatal(err)
	}
	return userName
}
*/
