package model

import (
	"context"
	"database/sql"
	"errors"
	"time"

	database "flover-market/db"
	"github.com/google/uuid"
)

/*CREATE TABLE IF NOT EXISTS sessions (
    session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    created_at INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    FOREIGN KEY ([customer_id]) REFERENCES "customers" ([customer_id]) ON DELETE NO ACTION ON UPDATE NO ACTION
);*/

type Session struct {
	SessionId  string `json:"session_id"`
	Created    int64  `json:"created"`
	CustomerId int64  `json:"customer_id"`
}

// CreateSession creates a new session for a customer. Returns the session ID.
func CreateSession(userId int64) (string, error) {
	db := database.GetDB()
	token := generateToken()
	_, err := db.Exec(context.Background(), "INSERT INTO sessions (session_id, created, customer_id) VALUES ($1, $2, $3)", token, time.Now().Unix(), userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func DeleteSession(token string) error {
	db := database.GetDB()
	_, err := db.Exec(context.Background(), "DELETE FROM sessions WHERE session_id = $1", token)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(token string) (Session, error) {
	db := database.GetDB()

	var session Session

	err := db.QueryRow(context.Background(), "SELECT customer_id, created, session_id FROM sessions WHERE session_id = $1", token).Scan(&session.CustomerId, &session.Created, &session.SessionId)
	if err != nil {
		return session, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return session, errors.New("session not found")
	}

	return session, nil
}

func generateToken() string {
	return uuid.New().String()
}
