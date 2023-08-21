package db

import (
	"fmt"

	"cs-skins-api/src/logger"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//market_hash_name VARCHAR(500) NOT NULL,
//currency VARCHAR(500) NOT NULL,
//suggested_price VARCHAR(500) NOT NULL,
//item_page VARCHAR(500) NOT NULL,
//market_page    VARCHAR(500) NOT NULL,
//max_price   VARCHAR(500)  NOT NULL,
//median_price    VARCHAR(500) NOT NULL,
//quantity   VARCHAR(500) NOT NULL,
//created_at   VARCHAR(500) NOT NULL,
//updated_at   VARCHAR(500) NOT NULL

type Request struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Wear        string `json:"wear"`
	SkinPattern string `json:"skin_pattern"`
	Price       string `json:"price"`
	Site        string `json:"site"`
}

type Skinport struct {
	ID             uint    `gorm:"primaryKey"`
	MarketHashName string  `json:"market_hash_name"`
	Currency       string  `json:"currency"`
	SuggestedPrice float64 `json:"suggested_price"`
	ItemPage       string  `json:"item_page"`
	MarketPage     string  `json:"market_page"`
	MaxPrice       float64 `json:"max_price"`
	MeanPrice      float64 `json:"mean_price"`
	Quantity       int     `json:"quantity"`
	CreatedAt      int     `json:"created_at"`
	UpdatedAt      int     `json:"updated_at"`
}

type Dmarket struct {
	ID      uint   `gorm:"primaryKey"`
	Cursor  string `json:"cursor"`
	Objects []struct {
		Amount      int    `json:"amount"`
		ClassID     string `json:"classId"`
		CreatedAt   int    `json:"createdAt"`
		Description string `json:"description"`
		Discount    int    `json:"discount"`
		Extra       struct {
			Ability         string   `json:"ability"`
			BackgroundColor string   `json:"backgroundColor"`
			Category        string   `json:"category"`
			CategoryPath    string   `json:"categoryPath"`
			Class           []string `json:"class"`
			Collection      []string `json:"collection"`
			Exterior        string   `json:"exterior"`
			FloatValue      int      `json:"floatValue"`
			GameID          string   `json:"gameId"`
			Gems            []struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Type  string `json:"type"`
			} `json:"gems"`
			Grade         string `json:"grade"`
			GroupID       string `json:"groupId"`
			Growth        int    `json:"growth"`
			Hero          string `json:"hero"`
			InspectInGame string `json:"inspectInGame"`
			IsNew         bool   `json:"isNew"`
			ItemType      string `json:"itemType"`
			LinkID        string `json:"linkId"`
			Name          string `json:"name"`
			NameColor     string `json:"nameColor"`
			OfferID       string `json:"offerId"`
			Quality       string `json:"quality"`
			Rarity        string `json:"rarity"`
			SerialNumber  int    `json:"serialNumber"`
			Stickers      []struct {
				Image string `json:"image"`
				Name  string `json:"name"`
			} `json:"stickers"`
			Subscribers       int    `json:"subscribers"`
			TagName           string `json:"tagName"`
			Tradable          bool   `json:"tradable"`
			TradeLock         int    `json:"tradeLock"`
			TradeLockDuration int    `json:"tradeLockDuration"`
			Type              string `json:"type"`
			Videos            int    `json:"videos"`
			ViewAtSteam       string `json:"viewAtSteam"`
			Withdrawable      bool   `json:"withdrawable"`
		} `json:"extra"`
		ExtraDoc     string `json:"extraDoc"`
		GameID       string `json:"gameId"`
		GameType     string `json:"gameType"`
		Image        string `json:"image"`
		InMarket     bool   `json:"inMarket"`
		InstantPrice struct {
			Dmc string `json:"DMC"`
			Usd string `json:"USD"`
		} `json:"instantPrice"`
		InstantTargetID string `json:"instantTargetId"`
		ItemID          string `json:"itemId"`
		LockStatus      bool   `json:"lockStatus"`
		Owner           string `json:"owner"`
		OwnerDetails    struct {
			Avatar string `json:"avatar"`
			ID     string `json:"id"`
			Wallet string `json:"wallet"`
		} `json:"ownerDetails"`
		OwnersBlockchainID string `json:"ownersBlockchainId"`
		Price              struct {
			Dmc string `json:"DMC"`
			Usd string `json:"USD"`
		} `json:"price"`
		RecommendedPrice struct {
			D3 struct {
				Dmc string `json:"DMC"`
				Usd string `json:"USD"`
			} `json:"d3"`
			D7 struct {
				Dmc string `json:"DMC"`
				Usd string `json:"USD"`
			} `json:"d7"`
			D7Plus struct {
				Dmc string `json:"DMC"`
				Usd string `json:"USD"`
			} `json:"d7Plus"`
		} `json:"recommendedPrice"`
		Slug           string `json:"slug"`
		Status         string `json:"status"`
		SuggestedPrice struct {
			Dmc string `json:"DMC"`
			Usd string `json:"USD"`
		} `json:"suggestedPrice"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"objects"`
	Total string `json:"total"`
}

type CsgoFloat struct {
	ID        uint      `gorm:"primaryKey"`
	CsFloatID string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	State     string    `json:"state"`
	Seller    struct {
		Avatar      string `json:"avatar"`
		Flags       int    `json:"flags"`
		Online      bool   `json:"online"`
		StallPublic bool   `json:"stall_public"`
		Statistics  struct {
			MedianTradeTime     int `json:"median_trade_time"`
			TotalFailedTrades   int `json:"total_failed_trades"`
			TotalTrades         int `json:"total_trades"`
			TotalVerifiedTrades int `json:"total_verified_trades"`
		} `json:"statistics"`
		SteamID  string `json:"steam_id"`
		Username string `json:"username"`
	} `json:"seller"`
	Item struct {
		AssetID        string  `json:"asset_id"`
		DefIndex       int     `json:"def_index"`
		PaintIndex     int     `json:"paint_index"`
		PaintSeed      int     `json:"paint_seed"`
		FloatValue     float64 `json:"float_value"`
		IconURL        string  `json:"icon_url"`
		DParam         string  `json:"d_param"`
		IsStattrak     bool    `json:"is_stattrak"`
		IsSouvenir     bool    `json:"is_souvenir"`
		Rarity         int     `json:"rarity"`
		Quality        int     `json:"quality"`
		MarketHashName string  `json:"market_hash_name"`
		Stickers       []struct {
			StickerID int    `json:"stickerId"`
			Slot      int    `json:"slot"`
			IconURL   string `json:"icon_url"`
			Name      string `json:"name"`
			Scm       struct {
				Price  int `json:"price"`
				Volume int `json:"volume"`
			} `json:"scm"`
		} `json:"stickers"`
		Tradable      int    `json:"tradable"`
		InspectLink   string `json:"inspect_link"`
		HasScreenshot bool   `json:"has_screenshot"`
		Scm           struct {
			Price  int `json:"price"`
			Volume int `json:"volume"`
		} `json:"scm"`
		ItemName    string `json:"item_name"`
		WearName    string `json:"wear_name"`
		Description string `json:"description"`
		Collection  string `json:"collection"`
		Badges      []any  `json:"badges"`
	} `json:"item"`
	IsSeller         bool `json:"is_seller"`
	MinOfferPrice    int  `json:"min_offer_price"`
	MaxOfferDiscount int  `json:"max_offer_discount"`
	IsWatchlisted    bool `json:"is_watchlisted"`
	Watchers         int  `json:"watchers"`
}

var requests []Request
var skinportSkins []Skinport
var csgoFloatSkins []CsgoFloat
var dmarketSkins []Dmarket

func ConnectToDb() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Babycakes15!@tcp(192.168.1.88:3306)/cs?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func WriteSkinportSkinToDb(skinport []Skinport, db *gorm.DB) {
	fmt.Println("Connected to DB.")

	if !db.Migrator().HasTable("skinport") {
		db.Migrator().CreateTable(&Skinport{})
		fmt.Println("Added table")
	}

	for i := 0; i < len(skinport); i++ {
		db.Select("MarketHashName", "Currency", "SuggestedPrice", "ItemPage", "MarketPage", "MaxPrice", "MeanPrice", "Quantity", "CreatedAt", "UpdatedAt").Create(&skinport[i])
	}
}

func WriteDmarketSkinToDb(dmarket []Dmarket, db *gorm.DB) {
	fmt.Println("Connected to DB.")

	if !db.Migrator().HasTable("dmarket") {
		db.Migrator().CreateTable(&Dmarket{})
		fmt.Println("Added table")
	}

	for i := 0; i < len(dmarket); i++ {
		db.Select("Amount", "ClassID", "Description", "Discount", "Total").Create(&dmarket[i])
	}
}

func WriteCsgoFloatSkinToDb(csgofloat []CsgoFloat, db *gorm.DB) {
	fmt.Println("Connected to DB.")

	if !db.Migrator().HasTable("csfloat") {
		db.Migrator().CreateTable(&CsgoFloat{})
		fmt.Println("Added table")
	}

	for i := 0; i < len(csgofloat); i++ {
		db.Select("CsFloatID", "CreatedAt", "Type", "Price", "State", "Seller", "Item", "IsSeller", "MinOfferPrice", "MaxOfferDiscount", "IsWatchlisted", "Watchers").Create(&csgofloat[i])
	}
}

func WriteRequestToDb(request *Request, db *gorm.DB) {
	fmt.Println("Connected to DB.")

	if !db.Migrator().HasTable("requests") {
		db.Migrator().CreateTable(&Request{})

		fmt.Println("Added table")
	}

	fmt.Println(request)

	db.Select("Name", "Wear", "SkinPattern", "Price").Create(&request)
}

func CompareRequestToSkinsDB(db *gorm.DB, element *Request) ([]string, []float64) {
	skins := []string{}
	prices := []float64{}

	fmt.Println("Connected to DB.")

	//    if db.Migrator().HasTable("requests") {
	//        _ = db.Table("requests").Select([]string{"name", "wear", "skin_pattern", "price"}).Scan(&requests)
	//
	//        fmt.Println(requests)
	//
	//        for _, element := range requests {
	_ = db.Table("skinport").Select([]string{"market_hash_name", "currency", "suggested_price", "item_page", "market_page", "max_price", "mean_price", "quantity", "created_at", "updated_at"}).Scan(&skinportSkins)
	_ = db.Table("csfloat").Select([]string{"type", "price", "state", "tradable", "item_name", "wear_name", "description", "collection", "badges", "market_hash_name"}).Scan(&csgoFloatSkins)
	_ = db.Table("dmarket").Select([]string{"amount", "classId", "description", "discount", "total"}).Scan(&dmarketSkins)

	for _, skinportSkin := range skinportSkins {
		price := element.Price

		// fmt.Println(price)

		requestedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			logger.WriteError("Could not parse requestedPrice. ", err)
		}

		// fmt.Println(requestedPrice)

		// splitString := strings.Split(skinportSkin.MarketHashName, " ")

		//                fmt.Println(skinportSkin.MarketHashName)
		//
		//                fmt.Println(requestedPrice, skinportSkin.MeanPrice)

		if strings.Contains(skinportSkin.MarketHashName, element.Wear) && strings.Contains(skinportSkin.MarketHashName, element.Name) && strings.Contains(skinportSkin.MarketHashName, element.SkinPattern) && skinportSkin.MeanPrice <= requestedPrice && skinportSkin.MeanPrice != 0 {
			skins = append(skins, skinportSkin.MarketHashName)
			prices = append(prices, skinportSkin.MeanPrice)
		}
	}

	for _, csgoFloatSkin := range csgoFloatSkins {
		price := element.Price

		requestedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			logger.WriteError("Could not parse requestedPrice. ", err)
		}

		// fmt.Println(requestedPrice)

		// splitString := strings.Split(skinportSkin.MarketHashName, " ")

		//                fmt.Println(skinportSkin.MarketHashName)
		//
		//                fmt.Println(requestedPrice, skinportSkin.MeanPrice)

		if strings.Contains(csgoFloatSkin.Item.MarketHashName, element.Wear) && strings.Contains(csgoFloatSkin.Item.MarketHashName, element.Name) && strings.Contains(csgoFloatSkin.Item.MarketHashName, element.SkinPattern) && float64(csgoFloatSkin.Price) <= requestedPrice && csgoFloatSkin.Price != 0 {
			skins = append(skins, csgoFloatSkin.Item.MarketHashName)
			prices = append(prices, float64(csgoFloatSkin.Price))
		}
	}
	//        }
	//    }

	return skins, prices
}
