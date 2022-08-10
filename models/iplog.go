package models

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type IpLog struct {
	ID        int64     `pg:"id, primarykey, autoincrement" json:"id"`
	Ip        string    `json:"ip"`
	Count     int       `json:"count"`
	SavedTime time.Time `json:"saved_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IpLogModel struct{}

// Create Address Table
func CreateIpLogTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&IpLog{}, opts)
	if createError != nil {
		log.Printf("Error while creating address table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("IpLog table created")
	return nil
}

func (m IpLogModel) CreateIpLog(address string) (*IpLog, error) {

	ipLogModel := &IpLog{
		Ip:        address,
		Count:     1,
		SavedTime: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	insertError := DbConnect.Insert(ipLogModel)
	if insertError != nil {
		return ipLogModel, insertError
	}
	return ipLogModel, nil
}

func (m IpLogModel) Select(id int64) (*IpLog, error) {
	ipLog := &IpLog{ID: id}
	addressErr := DbConnect.Select(ipLog)
	if addressErr != nil {
		return ipLog, addressErr
	}
	return ipLog, nil
}

func (m IpLogModel) SelectCustom(key string, value string) (*IpLog, error) {
	address := &IpLog{}
	err := DbConnect.Model(address).
		Where("? = ?", pg.Ident(key), value).
		Select()
	if err != nil {
		return address, err
	}
	return address, nil
}

func (m IpLogModel) SelectManyCustom(key string, value string) (*[]IpLog, error) {

	address := &[]IpLog{}
	err := DbConnect.Model(address).
		Where("? = ?", pg.Ident(key), value).
		Select()
	if err != nil {
		return address, err
	}
	return address, nil
}

func (m IpLogModel) OrCustom(key string, value string, secondKey string, secondValue string) (*IpLog, error) {

	address := &IpLog{}
	err := DbConnect.Model(address).
		Where("? = ?", pg.Ident(key), value).
		WhereOr("? = ?", pg.Ident(secondKey), secondValue).
		Select()
	if err != nil {
		return address, err
	}
	return address, nil
}

func (m IpLogModel) Update(ipLog *IpLog) (*IpLog, error) {

	_, err := DbConnect.Model(ipLog).
		WherePK().
		UpdateNotZero()
	if err != nil {
		return ipLog, err
	}
	return ipLog, nil
}
