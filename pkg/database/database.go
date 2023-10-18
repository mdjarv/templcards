package database

import (
	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/mdjarv/templcards/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	db *gorm.DB
}

func New() (*Database, error) {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (db Database) Migrate() error {
	return db.db.AutoMigrate(&models.Card{})
}

func (db Database) GetCards() ([]models.Card, error) {
	var cards []models.Card
	result := db.db.Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}
	return cards, nil
}

func (db Database) AddCard(card models.Card) (models.Card, error) {
	card.ID = uuid.New().String()
	err := db.db.Clauses(clause.Returning{}).Create(&card).Error

	return card, err
}
