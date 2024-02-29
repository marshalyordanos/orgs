package database

import (
	"database/sql"
	"fmt"

	"auth/config"

	_ "github.com/lib/pq"
)

type postgresDatabase struct {
	Db *sql.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) GetDb() *sql.DB {
	return p.Db
}
