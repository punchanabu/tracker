// interface/api/handler/wallet.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/entity"
	"github.com/punchanabu/portfolio-tracker/internal/domain/service"
	"github.com/punchanabu/portfolio-tracker/internal/domain/vo"
)

type WalletHandler struct {
	walletService *service.WalletService
}

func NewWalletHandler(walletService *service.WalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

type AddWalletRequest struct {
	PortfolioID string `json:"portfolioId" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Chain       string `json:"chain" binding:"required"`
}

func (h *WalletHandler) AddWallet(c *gin.Context) {
	var req AddWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	portfolioID, err := uuid.Parse(req.PortfolioID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid portfolio ID"})
		return
	}

	address := vo.NewAddress(req.Address)
	if !address.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet address"})
		return
	}

	wallet := &entity.Wallet{
		Address: address,
		Chain:   req.Chain,
	}

	err = h.walletService.AddWalletToPortfolio(c.Request.Context(), portfolioID, wallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

func (h *WalletHandler) GetWalletBalance(c *gin.Context) {
	walletID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid wallet ID"})
		return
	}

	balance, err := h.walletService.GetWalletBalance(c.Request.Context(), walletID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
