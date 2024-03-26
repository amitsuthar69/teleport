package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	encode "teleport/internal/utils"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Service interface {
	Health() map[string]string
	SetLongUrl(string) string
	GetLongUrl(string) string
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

func (s *service) SetLongUrl(longUrl string) string {
	time := time.Now().Format("06121545")
	id, err := strconv.Atoi(time)
	if err != nil {
		fmt.Println(err)
	}
	hash := encode.Base62(int64(id))
	stmt := `INSERT INTO shortUrls (shortUrl, longUrl) VALUES (?, ?)`
	_, err = s.db.Exec(stmt, hash, longUrl)
	if err != nil {
		log.Printf("Error inserting short URL: %v", err)
		return err.Error()
	}
	return hash
}

func (s *service) GetLongUrl(hash string) string {
	stmt := `SELECT longUrl FROM shortUrls WHERE shortUrl= ? `
	row := s.db.QueryRow(stmt, hash)
	var LongUrl string
	err := row.Scan(&LongUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows selected")
			return ""
		}
		log.Fatal(err)
	}
	return LongUrl
}
