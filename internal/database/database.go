package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"teleport/internal/cache"
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
	db    *sql.DB
	cache *cache.CacheMap
}

var (
	dburl = os.Getenv("DB_URL")
)

func New() Service {
	db, err := sql.Open("libsql", dburl)
	if err != nil {
		log.Fatal(err)
	}
	s := &service{db: db, cache: cache.NewCacheMap(50)}
	return s
}

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

func (s *service) SetLongUrl(longUrl string) string {
	var existingHash string
	checkStmt := `SELECT shortUrl FROM shortUrls WHERE longUrl = ?`
	err := s.db.QueryRow(checkStmt, longUrl).Scan(&existingHash)
	if err == nil {
		s.cache.Set(existingHash, longUrl)
		return existingHash
	} else if err != sql.ErrNoRows {
		log.Printf("Error checking existing long URL: %v", err)
		return err.Error()
	}

	id := time.Now().UnixNano()
	hash := encode.Base62(int64(id))
	stmt := `INSERT INTO shortUrls (shortUrl, longUrl) VALUES (?, ?)`
	_, err = s.db.Exec(stmt, hash, longUrl)
	if err != nil {
		log.Printf("Error inserting short URL: %v", err)
		return err.Error()
	}

	s.cache.Set(hash, longUrl)
	return hash
}

func (s *service) GetLongUrl(hash string) string {
	if url, exists := s.cache.Get(hash); exists {
		log.Printf("Cache hit for hash: %s", hash)
		return url
	}
	log.Printf("Cache miss for hash: %s", hash)

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

	s.cache.Set(hash, LongUrl)
	return LongUrl
}
