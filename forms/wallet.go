package forms

//WalletForm ...
type WalletForm struct{}

//AccountForm ...
type AccountForm struct {
	PublicKey string `form:"publicKey" json:"publicKey" binding:"required"`
	SignAlgo  int    `form:"signatureAlgorithm" json:"signatureAlgorithm" binding:"required"`
	HashAlgo  int    `form:"hashAlgorithm" json:"hashAlgorithm" binding:"required"`
}
