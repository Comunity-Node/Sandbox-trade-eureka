package routes

import (
    "github.com/gin-gonic/gin"
    "sandbox-invest/handlers"
    "sandbox-invest/auth"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/portfolio/:userId", handlers.GetPortfolio)
    r.GET("/prices", handlers.GetPrices)
    r.GET("/price", handlers.GetLivePrice)

    authGroup := r.Group("/")
    authGroup.Use(auth.AuthRequired())
    authGroup.POST("/buy", handlers.BuyAsset)
    authGroup.POST("/sell", handlers.SellAsset)

    return r
}
