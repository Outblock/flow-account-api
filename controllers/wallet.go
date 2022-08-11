package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onflow/flow-go-sdk"
	"outblock.io/go-server/demo/config"
	"outblock.io/go-server/demo/forms"
	models "outblock.io/go-server/demo/models"
)

//WalletController ...
type WalletController struct{}

var walletModel = new(models.WalletModel)
var walletModelMain = new(models.WalletModelMain)
var ipLogModel = new(models.IpLogModel)

type WalletReturn struct {
	TxId string `json:"txId"`
}

// Get accounts with public key godoc
// @Summary      Get accounts
// @Description  Submit public key to receive accounts
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {array} []models.Wallet "return 200 with the array of accounts."
// @Router       /v1/address [get]
func (ctrl WalletController) Getrecord(c *gin.Context) {

	publicKey := c.Query("publicKey")

	wallet, walletErr := walletModelMain.SelectManyCustom("public_key", publicKey)
	if walletErr != nil {
		log.Printf("Error while creating a wallet, Reason: %v\n", walletErr)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Request error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Request success",
		"data":    wallet,
	})

	return
}

// Get accounts with public key testnet
// @Summary      Get accounts
// @Description  Submit public key to receive accounts
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {array} []models.Wallet "return 200 with the array of accounts from testnet."
// @Router       /v1/address/testnet [get]
func (ctrl WalletController) GetrecordTest(c *gin.Context) {

	publicKey := c.Query("publicKey")

	wallet, walletErr := walletModel.SelectManyCustom("public_key", publicKey)
	if walletErr != nil {
		log.Printf("Error while creating a wallet, Reason: %v\n", walletErr)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Request error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Request success",
		"data":    wallet,
	})

	return
}

// Create address godoc
// @Summary      Create an address
// @Description  create address use public key
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        publicKey    body     string  true  "public key"
// @Param        signatureAlgorithm    body     string  true  "sign algorithm"
// @Param        hashAlgorithm    body     string  true  "hash algorithm"
// @Param        weight    body     int  true  "weight of the key"
// @Success      200  {object}  WalletReturn "return 200 with the transaction id."
// @Router       /v1/address [post]
func (ctrl WalletController) CreateAddress(c *gin.Context) {
	// check registered ip
	// ip := c.ClientIP()
	// ipLog, ipLogErr := ipLogModel.SelectCustom("ip", ip)
	// if ipLogErr != nil && ipLog.ID == 0 {
	// 	ipLog, _ = ipLogModel.CreateIpLog(ip)
	// } else {
	// 	if ipLog.SavedTime.After(time.Now().Add(-(time.Minute * 10))) && ipLog.Count > 5 {
	// 		fmt.Printf("The HTTP request failed with error")
	// 		c.JSON(429, gin.H{
	// 			"status":  429,
	// 			"message": "Temporary registration limit exceeded, please try again later",
	// 		})
	// 		return
	// 	} else if ipLog.SavedTime.Before(time.Now().Add(-(time.Minute * 10))) {
	// 		ipLog.SavedTime = time.Now()
	// 		ipLog.Count = 1
	// 	} else {
	// 		count := ipLog.Count
	// 		ipLog.Count = count + 1
	// 	}
	// 	ipLogModel.Update(ipLog)
	// }
	var accountForm forms.AccountForm
	if validationErr := c.ShouldBindJSON(&accountForm); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "validation error"})
		return
	}
	walletId, walletErr := walletModelMain.CreateWallet(accountForm)

	wallet, _ := walletModelMain.Select(walletId)

	if walletErr != nil {
		log.Printf("Error while creating a wallet, Reason: %v\n", walletErr)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Transaction error",
		})
		return
	}

	txId := createAccountMain(wallet, c)
	result := txId.String()

	var returnStruct WalletReturn
	returnStruct.TxId = result

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Transaction created",
		"data":    returnStruct,
	})

	return
}

//Create flow account
func createAccountMain(wallet *models.WalletMain, c *gin.Context) flow.Identifier {

	txMain := config.CreateFlowKey(wallet.HashAlgoString, wallet.SignAlgoString, wallet.PublicKey, wallet.Weight, "mainnet", c)
	txidMain := txMain.ID()

	go generateWalletMain(txidMain, wallet)

	return txidMain
}

//Receive and save flow account mainnet
func generateWalletMain(id flow.Identifier, wallet *models.WalletMain) error {
	result := config.WaitForSeal(id, "mainnet")

	err := saveWalletMain(wallet, result)
	if err != nil {
		log.Printf("Error while generating, Reason: %v\n", err)
		return err
	}

	return nil
}

//Save the flow account to database
func saveWalletMain(wallet *models.WalletMain, result string) error {

	savedAddress := "0x" + result
	wallet.Address = savedAddress
	_, walletErr := walletModelMain.Update(wallet)

	if walletErr != nil {
		log.Printf("Error while saving, Reason: %v\n", walletErr)
		return walletErr
	}

	return nil
}

// Create address on testnet
// @Summary      Create an address
// @Description  use public key to create address
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        publicKey    body     string  true  "public key"
// @Param        signatureAlgorithm    body     string  true  "sign algorithm"
// @Param        hashAlgorithm    body     string  true  "hash algorithm"
// @Param        weight    body     int  true  "weight of the key"
// @Success      200   {object}  WalletReturn "return 200 with the transaction id."
// @Router       /v1/address/testnet [post]
func (ctrl WalletController) CreateAddressTest(c *gin.Context) {
	// check registered ip
	// ip := c.ClientIP()
	// ipLog, ipLogErr := ipLogModel.SelectCustom("ip", ip)
	// if ipLogErr != nil && ipLog.ID == 0 {
	// 	ipLog, _ = ipLogModel.CreateIpLog(ip)
	// } else {
	// 	if ipLog.SavedTime.After(time.Now().Add(-(time.Minute * 10))) && ipLog.Count > 5 {
	// 		fmt.Printf("The HTTP request failed with error")
	// 		c.JSON(429, gin.H{
	// 			"status":  429,
	// 			"message": "Temporary registration limit exceeded, please try again later",
	// 		})
	// 		return
	// 	} else if ipLog.SavedTime.Before(time.Now().Add(-(time.Minute * 10))) {
	// 		ipLog.SavedTime = time.Now()
	// 		ipLog.Count = 1
	// 	} else {
	// 		count := ipLog.Count
	// 		ipLog.Count = count + 1
	// 	}
	// 	ipLogModel.Update(ipLog)
	// }
	var accountForm forms.AccountForm

	if validationErr := c.ShouldBindJSON(&accountForm); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "validation error"})
		return
	}
	walletId, walletErr := walletModel.CreateWallet(accountForm)

	wallet, _ := walletModel.Select(walletId)
	if walletErr != nil {
		log.Printf("Error while creating a wallet, Reason: %v\n", walletErr)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Transaction error",
		})
		return
	}

	txId := createAccountTest(wallet, c)
	result := txId.String()

	var returnStruct WalletReturn
	returnStruct.TxId = result

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Transaction created",
		"data":    returnStruct,
	})

	return
}

func createAccountTest(wallet *models.Wallet, c *gin.Context) flow.Identifier {

	tx := config.CreateFlowKey(wallet.HashAlgoString, wallet.SignAlgoString, wallet.PublicKey, wallet.Weight, "testnet", c)
	txid := tx.ID()

	go generateWalletTest(txid, wallet)

	return txid
}

//Receive and save flow account
func generateWalletTest(id flow.Identifier, wallet *models.Wallet) error {
	result := config.WaitForSeal(id, "testnet")

	err := saveWalletTest(wallet, result)
	if err != nil {
		log.Printf("Error while creating a wallet, Reason: %v\n", err)
		return err
	}

	return nil
}

//Save the flow account to database
func saveWalletTest(wallet *models.Wallet, result string) error {

	savedAddress := "0x" + result
	wallet.Address = savedAddress
	_, walletErr := walletModel.Update(wallet)

	if walletErr != nil {
		log.Printf("Error while creating wallet address relation, Reason: %v\n", walletErr)
		return walletErr
	}

	return nil
}
