package database

import (
	"database/sql"
	"fmt"
	"kodarsiv/chatGPT-slack/internal/storage"
	"kodarsiv/chatGPT-slack/internal/types"
)

type Mysql struct {
	db *sql.DB
}

type Credentials struct {
	Host     string
	Username string
	Password string
	Database string
	Port     int
}

func NewMysql(credentials Credentials) (storage.Storage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", credentials.Username, credentials.Password, credentials.Host, credentials.Port, credentials.Database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	return &Mysql{db: db}, nil
}

func (m *Mysql) PutMessage(message types.Message) (bool, error) {
	// todo:: will come
	return false, nil
}
func (m *Mysql) Close() error {
	if err := m.db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %v", err)
	}
	return nil
}
