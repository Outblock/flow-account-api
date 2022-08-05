package forms

//WalletForm ...
type WalletForm struct{}

//AccountForm ...
type AccountForm struct {
	PublicKey string `form:"public_key" json:"public_key" binding:"required"`
	SignAlgo  string `form:"sign_algo" json:"sign_algo" binding:"required"`
	HashAlgo  string `form:"hash_algo" json:"hash_algo" binding:"required"`
	Weight    int    `form:"weight" json:"weight" binding:"required"`
}
