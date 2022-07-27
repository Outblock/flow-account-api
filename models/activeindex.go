package models

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type Activeindex struct {
	ID        int64     `pg:"id, primarykey, autoincrement" json:"id"`
	Network   string    `json:"network" default:"0"`
	Index     int64     `json:"index" default:"0"`
	MaxIndex  int64     `json:"maxIndex" default:"0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ActiveindexModel struct{}

// Create Address Table
func CreateActiveindexTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Activeindex{}, opts)
	if createError != nil {
		log.Printf("Error while creating address table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Address table created")
	return nil
}

func (m ActiveindexModel) CreateActiveindex(network string, maxIndex int64) (int64, error) {

	indexModel := &Activeindex{
		Network:  network,
		Index:    0,
		MaxIndex: maxIndex,
	}
	insertError := DbConnect.Insert(indexModel)
	if insertError != nil {
		return 0, insertError
	}
	return indexModel.ID, nil
}

func (m ActiveindexModel) Select(id int64) (*Activeindex, error) {
	activeIndex := &Activeindex{ID: id}
	err := DbConnect.Select(activeIndex)
	if err != nil {
		return activeIndex, err
	}
	return activeIndex, nil
}

func (m ActiveindexModel) SelectCustom(key string, value string) (*Activeindex, error) {
	activeIndex := &Activeindex{}
	err := DbConnect.Model(activeIndex).
		Where("? = ?", pg.Ident(key), value).
		Select()
	if err != nil {
		return activeIndex, err
	}
	return activeIndex, nil
}

func (m ActiveindexModel) Update(activeIndex *Activeindex) (*Activeindex, error) {

	_, err := DbConnect.Model(activeIndex).
		WherePK().
		Update()
	if err != nil {
		return activeIndex, err
	}
	return activeIndex, nil
}
