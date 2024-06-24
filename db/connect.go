package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func Connect(url string) error {
	var err error
	db, err = pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}
	return nil
}

// Init загружаем схему в базу
func Init() error {
	err := db.Ping(context.Background())
	if err != nil {
		return err
	}
	_, err = db.Exec(context.Background(), script)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *pgxpool.Pool {
	return db
}
