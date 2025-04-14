package postgres

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/FrancoBarrera99/auth-service/internal/auth/model"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserMigration struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	Username  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

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

	if err := db.AutoMigrate(&UserMigration{}); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) CreateUser(username string, password string, email string) (*model.User, error) {
	// Validate parameters
	if username == "" || password == "" || email == "" {
		return nil, fmt.Errorf("all fields are required")
	}
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("invalid email format")
	}

	// Create user
	user := model.User{
		ID:       uuid.New().String(),
		Username: username,
		Password: password,
		Email:    email,
	}

	if createErr := p.db.Create(&user).Error; createErr != nil {
		return nil, createErr
	}

	return &user, nil
}

func (p *PostgresStorage) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
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
