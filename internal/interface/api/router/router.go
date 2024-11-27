// interface/api/router/router.go
package router

import (
	"github.com/punchanabu/portfolio-tracker/internal/interface/api/handler"
	"github.com/punchanabu/portfolio-tracker/internal/interface/api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	portfolioHandler *handler.PortfolioHandler,
	walletHandler *handler.WalletHandler,
) *gin.Engine {
	router := gin.New()

	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.LoggerMiddleware())

	public := router.Group("/api/v1")
	{
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		portfolios := protected.Group("/portfolios")
		{
			portfolios.POST("/", portfolioHandler.Create)
			portfolios.GET("/:id", portfolioHandler.GetByID)
			portfolios.POST("/:id/wallets", walletHandler.AddWallet)
		}

		wallets := protected.Group("/wallets")
		{
			wallets.GET("/:id/balance", walletHandler.GetWalletBalance)
		}
	}

	return router
}
