package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/punchanabu/portfolio-tracker/internal/domain/service"
)

type PortfolioHandler struct {
	portfolioService *service.PortfolioService
}

func NewPortfolioHandler(portfolioService *service.PortfolioService) *PortfolioHandler {
	return &PortfolioHandler{
		portfolioService: portfolioService,
	}
}

type CreatePortfolioRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

func (h *PortfolioHandler) Create(c *gin.Context) {
	var req CreatePortfolioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	portfolio, err := h.portfolioService.CreatePortfolio(c.Request.Context(), userID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, portfolio)
}

func (h *PortfolioHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid portfolio ID"})
		return
	}

	portfolio, err := h.portfolioService.GetPortfolioByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

func (h *PortfolioHandler) GetByUserID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	portfolios, err := h.portfolioService.GetPortfoliosByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, portfolios)
}
