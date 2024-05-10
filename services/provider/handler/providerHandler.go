package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/services/provider/gen"
)

type ProviderHandler struct {
	db ProviderDatabase
}

type ProviderDatabase interface {
	GetProvider(id string) (*gen.Provider, error)
	GetProviders(params gen.GetProvidersParams) ([]gen.Provider, error)
	CreateProvider(provider *gen.NewProvider) (*gen.Provider, error)
	DeleteProvider(id string) error
}

func NewProviderHandler(db ProviderDatabase) *ProviderHandler {
	return &ProviderHandler{db: db}
}

func (h *ProviderHandler) CheckHealth(c *gin.Context) {
	c.Status(200)
}

func (h *ProviderHandler) GetProvider(c *gin.Context, id string) {
	provider, err := h.db.GetProvider(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, provider)
}

func (h *ProviderHandler) GetProviders(c *gin.Context, params gen.GetProvidersParams) {
	providers, err := h.db.GetProviders(params)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, providers)
}

func (h *ProviderHandler) CreateProvider(c *gin.Context) {
	var provider gen.NewProvider
	if err := c.BindJSON(&provider); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(provider)

	createdProvider, err := h.db.CreateProvider(&provider)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdProvider)
}

func (h *ProviderHandler) DeleteProvider(c *gin.Context, id string) {
	err := h.db.DeleteProvider(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
