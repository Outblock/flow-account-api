package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"outblock.io/go-server/demo/models"
)

var activeindexModel = new(models.ActiveindexModel)

func Flow(network string) string {
	switch network {
	case "testnet":
		return "access.devnet.nodes.onflow.org:9000"
	case "mainnet":
		return "access.mainnet.nodes.onflow.org:9000"
	default:
		return "access.devnet.nodes.onflow.org:9000"
	}
}


func FlowKey(network string) (string, string, int) {
	switch network {
	case "testnet":
		return "0x331c3648941e4863", "0e3d8dbae1f97ed41fa39ba797fc9eed0dae80a6b630ffbb32b2b389386ea29c", 0
	case "mainnet":
		return "0x46a783bae80ddeb2", "0e3d8dbae1f97ed41fa39ba797fc9eed0dae80a6b630ffbb32b2b389386ea29c", 0

		// return "33f75ff0b830dcec", "b9e452a92cdfc4857b6b6dd499f28e310c314818b7fd5f05c34736c0a7eb1255"
	default:
		return "0x331c3648941e4863", "0e3d8dbae1f97ed41fa39ba797fc9eed0dae80a6b630ffbb32b2b389386ea29c", 0
	}
}

func GetKey(network string, c *gin.Context) (string, string, int64) {
	activeIndex, activeIndexErr := activeindexModel.SelectCustom("network", network)

	if activeIndexErr != nil {
		log.Println("active index error: ", activeIndexErr)
	}

	currentIndex := activeIndex.Index
	address, privateKey, _ := FlowKey(network)
	//return value through firestore

	maxIndex := activeIndex.MaxIndex

	activeIndex.Index = currentIndex + 1

	if currentIndex == maxIndex {
		activeIndex.Index = 1
	}
	_, updateErr := activeindexModel.Update(activeIndex)

	if updateErr != nil {
		log.Println(updateErr)
	}

	return address, privateKey, currentIndex
}

func Gdrive() string {
	return "https://www.googleapis.com/drive/v3/files/"
}
