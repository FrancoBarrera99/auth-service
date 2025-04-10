package postgres

import (
	"fmt"
	"os"

	m "github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, fmt.Errorf("invalid DSN value")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) SaveUser(user m.User) error {
	return p.db.Save(&user).Error
}

func (p *PostgresStorage) GetUserByEmail(email string) (*m.User, error) {
	var user m.User
	err := p.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresStorage) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return sqlDB.Close()
	}
	return err
}
