package models

import (
	"log"
	"time"

	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
	"outblock.io/go-server/demo/forms"
)

type Wallet struct {
	ID             int64     `pg:"id, primarykey, autoincrement" json:"id"`
	PublicKey      string    `json:"public_key"`
	Address        string    `json:"address"`
	HashAlgo       int       `json:"hash_algo"`
	SignAlgo       int       `json:"sign_algo"`
	Weight         int       `json:"weight"`
	HashAlgoString string    `json:"hash_algo_string"`
	SignAlgoString string    `json:"sign_algo_string"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WalletModel struct{}

// Create Wallet Table
func CreateWalletTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Wallet{}, opts)
	if createError != nil {
		log.Printf("Error while creating wallet table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Wallet table created")
	return nil
}

func (m WalletModel) CreateWallet(reqData forms.AccountForm) (int64, error) {

	walletModel := &Wallet{
		PublicKey:      reqData.PublicKey,
		HashAlgoString: reqData.HashAlgo,
		SignAlgoString: reqData.SignAlgo,
		Weight:         reqData.Weight,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	insertError := DbConnect.Insert(walletModel)

	if insertError != nil {
		return 0, insertError
	}
	return walletModel.ID, nil
}

func (m WalletModel) Select(id int64) (*Wallet, error) {
	wallet := &Wallet{ID: id}
	walletErr := DbConnect.Select(wallet)
	if walletErr != nil {
		return wallet, walletErr
	}
	return wallet, nil
}

func (m WalletModel) SelectCustom(key string, value string) (*Wallet, error) {

	wallet := &Wallet{}
	err := DbConnect.Model(wallet).
		Where("? = ?", pg.Ident(key), value).
		Select()
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (m WalletModel) SelectManyCustom(key string, value string) (*[]Wallet, error) {

	wallet := &[]Wallet{}
	err := DbConnect.Model(wallet).
		Where("? = ?", pg.Ident(key), value).
		Select()
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (m WalletModel) Update(wallet *Wallet) (*Wallet, error) {

	_, err := DbConnect.Model(wallet).
		WherePK().
		Update()
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}
