package forms

//WalletForm ...
type WalletForm struct{}

//AccountForm ...
type AccountForm struct {
	PublicKey string `form:"publicKey" json:"publicKey" binding:"required"`
	SignAlgo  string `form:"signatureAlgorithm" json:"signatureAlgorithm" binding:"required"`
	HashAlgo  string `form:"hashAlgorithm" json:"hashAlgorithm" binding:"required"`
	Weight    int    `form:"weight" json:"weight" binding:"required"`
}
