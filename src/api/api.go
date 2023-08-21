package api

import (
	"cs-skins-api/src/db"
	"fmt"
	"net/http"

	"cs-skins-api/src/logger"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//type Skinport struct {
//    ID     string  `json:"id"`
//    MarketHashName        string `json:"market_hash_name"`
//    Currency        string `json:"currency"`
//    SuggestedPrice string `json:"suggested_price"`
//    ItemPage       string `json:"item_page"`
//    MarketPage       string `json:"market_page"`
//    MaxPrice       string `json:"max_price"`
//    MedianPrice       string `json:"median_price"`
//    Quantity       string `json:"quantity"`
//    CreatedAt       string `json:"created_at"`
//    UpdatedAt       string `json:"updated_at"`
//}

func AddSkinportSkin(c *gin.Context) {
	dbConnected, err := db.ConnectToDb()
	if err != nil {
		logger.WriteError("Can't connect to the DB", err)
	}

	respBody := c.Request.Body

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		logger.WriteError("Failed to read the response body.", err)
	}

	var skinport []db.Skinport

	err = json.Unmarshal(body, &skinport)
	if err != nil {
		logger.WriteError("Failed to Unmarshal JSON.", err)
	}

	db.WriteSkinportSkinToDb(skinport, dbConnected)

	logger.WriteCMDInfo("go run main.go", "Success")

	c.JSON(http.StatusOK, gin.H{
		"message": "Skinport Update was successfully run.",
	})
}

func AddCSFloatSkin(c *gin.Context) {
	dbConnected, err := db.ConnectToDb()
	if err != nil {
		logger.WriteError("Can't connect to the DB", err)
	}

	respBody := c.Request.Body

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		logger.WriteError("Failed to read the response body.", err)
	}

	var csFloat []db.CsgoFloat

	err = json.Unmarshal(body, &csFloat)
	if err != nil {
		logger.WriteError("Failed to Unmarshal JSON.", err)
	}

	db.WriteCsgoFloatSkinToDb(csFloat, dbConnected)

	logger.WriteCMDInfo("go run main.go", "Success")

	c.JSON(http.StatusOK, gin.H{
		"message": "Skinport Update was successfully run.",
	})
}

func AddDMarketSkin(c *gin.Context) {
	dbConnected, err := db.ConnectToDb()
	if err != nil {
		logger.WriteError("Can't connect to the DB", err)
	}

	respBody := c.Request.Body

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		logger.WriteError("Failed to read the response body.", err)
	}

	var dMarket []db.Dmarket

	err = json.Unmarshal(body, &dMarket)
	if err != nil {
		logger.WriteError("Failed to Unmarshal JSON.", err)
	}

	db.WriteDmarketSkinToDb(dMarket, dbConnected)

	logger.WriteCMDInfo("go run main.go", "Success")

	c.JSON(http.StatusOK, gin.H{
		"message": "Skinport Update was successfully run.",
	})
}

func AddRequest(c *gin.Context) {
	dbConnected, err := db.ConnectToDb()
	if err != nil {
		logger.WriteError("Can't connect to the DB", err)
	}

	respBody := c.Request.Body

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		logger.WriteError("Failed to read the response body.", err)
	}

	var request *db.Request

	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.WriteError("Failed to Unmarshal JSON.", err)
	}

	fmt.Println(string(body))

	db.WriteRequestToDb(request, dbConnected)

	logger.WriteCMDInfo("go run main.go", "Success")

	c.JSON(http.StatusOK, gin.H{
		"message": "Skinport Update was successfully run.",
	})
}

func CompareRequestToSkins(c *gin.Context) {
	dbConnected, err := db.ConnectToDb()
	if err != nil {
		logger.WriteError("Can't connect to the DB", err)
	}

	respBody := c.Request.Body

	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		logger.WriteError("Failed to read the response body.", err)
	}

	var request *db.Request

	err = json.Unmarshal(body, &request)
	if err != nil {
		logger.WriteError("Failed to Unmarshal JSON.", err)
	}

	fmt.Println(string(body))

	skins, prices := db.CompareRequestToSkinsDB(dbConnected, request)

	c.JSON(http.StatusOK, gin.H{
		"skins":  skins,
		"prices": prices,
	})
}
