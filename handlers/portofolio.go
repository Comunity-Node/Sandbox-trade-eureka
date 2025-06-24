package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "sandbox-invest/models"
    "sandbox-invest/data"
    "sandbox-invest/services"
    "sandbox-invest/auth"
)

func GetPortfolio(c *gin.Context) {
    userId := c.Param("userId")
    p, exists := data.Portfolios[userId]
    if !exists {
        p = &models.Portfolio{Holdings: map[string]float64{}}
        data.Portfolios[userId] = p
    }
    c.JSON(http.StatusOK, p)
}

func GetPrices(c *gin.Context) {
    livePrice, err := services.GetYahooPrice("BBCA")
    if err == nil {
        data.Prices["BBCA"] = models.Asset{
            Code: "BBCA", Name: "Bank BCA", Type: models.Saham, Price: livePrice.RegularMarketPrice,
        }
    }

    prices := []models.Asset{}
    for _, asset := range data.Prices {
        prices = append(prices, asset)
    }
    c.JSON(http.StatusOK, prices)
}

func GetLivePrice(c *gin.Context) {
    symbol := c.Query("symbol")
    if symbol == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "symbol required"})
        return
    }

    priceData, err := services.GetYahooPrice(symbol)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, priceData)
}

func BuyAsset(c *gin.Context) {
    var tx models.Transaction
    if err := c.ShouldBindJSON(&tx); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    realUser := c.GetString("userId")
    if realUser != tx.UserID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "token does not match user_id"})
        return
    }

    custodyID := auth.GetCustodyUser(tx.UserID)
    if _, ok := data.Portfolios[custodyID]; !ok {
        data.Portfolios[custodyID] = &models.Portfolio{Holdings: map[string]float64{}}
    }

    data.Portfolios[custodyID].Holdings[tx.Code] += tx.Amount
    c.JSON(http.StatusOK, gin.H{"message": "buy success", "custody": custodyID})
}

func SellAsset(c *gin.Context) {
    var tx models.Transaction
    if err := c.ShouldBindJSON(&tx); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    realUser := c.GetString("userId")
    if realUser != tx.UserID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "token does not match user_id"})
        return
    }

    custodyID := auth.GetCustodyUser(tx.UserID)
    p, exists := data.Portfolios[custodyID]
    if !exists || p.Holdings[tx.Code] < tx.Amount {
        c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient holdings"})
        return
    }

    p.Holdings[tx.Code] -= tx.Amount
    c.JSON(http.StatusOK, gin.H{"message": "sell success", "custody": custodyID})
}
