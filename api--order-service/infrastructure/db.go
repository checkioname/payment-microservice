package infrastructure

import (
	"anturiocode/api--order-service/infrastructure/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStore interface {
	ExecQuery(query string, args ...interface{}) error
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

type postgresStore struct {
	db *sql.DB
}

func NewPostgresStore(cfg config.DatabaseConfig) (PostgresStore, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Host, cfg.Port, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &postgresStore{db: db}, nil

}

func (p *postgresStore) Close() error {
	return p.db.Close()
}

func (p *postgresStore) ExecQuery(query string, arg ...interface{}) error {
	_, err := p.db.Exec(query)
	if err != nil {
		fmt.Println("Não foi possível consultar o banco", err)
	}
	return err
}

func (p *postgresStore) QueryRow(query string, args ...interface{}) *sql.Row {
	return p.db.QueryRow(query, args...)
}

func (p *postgresStore) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.Query(query, args)
}
