package repository

import (
	"errors"
	"fmt"
	"gohex1/internal/core/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func NewMessengerPostgresRepository() *MessengerPostgresRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := ""
	dbname := "gohex"

	// "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
	// conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
	// 	host,
	// 	port,
	// 	user,
	// 	dbname,
	// 	password,
	// )

	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Message{})

	return &MessengerPostgresRepository{
		db: db,
	}

}

func (m *MessengerPostgresRepository) SaveMessage(message domain.Message) error {
	req := m.db.Create(&message)
	if req.RowsAffected == 0 {
		return fmt.Errorf("messages not saved: %v", req.Error)
	}
	return nil
}

func (m *MessengerPostgresRepository) ReadMessage(id string) (*domain.Message, error) {
	message := &domain.Message{}
	req := m.db.First(&message, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("message not found")
	}
	return message, nil
}

func (m *MessengerPostgresRepository) ReadMessages() ([]*domain.Message, error) {
	var messages []*domain.Message
	req := m.db.Find(&messages)
	if req.Error != nil {
		return nil, fmt.Errorf("messages not found: %v", req.Error)
	}
	return messages, nil
}
