package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/services/provider/gen"
)

// ProviderHandler defines the handlers for the Provider service
type ProviderHandler struct {
	db ProviderDatabase
}

// ProviderDatabase defines the database operations required for the Provider service
type ProviderDatabase interface {
	GetProvider(id string) (*gen.Provider, error)
	GetProviders(params gen.GetProvidersParams) ([]gen.Provider, error)
	CreateProvider(provider *gen.NewProvider) (*gen.Provider, error)
	DeleteProvider(id string) error
}

// NewProviderHandler creates a new ProviderHandler. It requires a ProviderDatabase
func NewProviderHandler(db ProviderDatabase) *ProviderHandler {
	return &ProviderHandler{db: db}
}

func (h *ProviderHandler) CheckHealth(c *gin.Context) {
	c.Status(200)
}

// GetProvider calls ProviderDatabase.GetProvider() and returns the result. It expects an id parameter
func (h *ProviderHandler) GetProvider(c *gin.Context, id string) {
	provider, err := h.db.GetProvider(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, provider)
}

// GetProviders calls ProviderDatabase.GetProviders() and returns the result. It accepts query parameters.
func (h *ProviderHandler) GetProviders(c *gin.Context, params gen.GetProvidersParams) {
	providers, err := h.db.GetProviders(params)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, providers)
}

// CreateProvider calls ProviderDatabase.CreateProvider() and returns the result. It expects a gen.NewProvider in the request body
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

// DeleteProvider calls ProviderDatabase.DeleteProvider() and returns the result. It expects an id parameter
func (h *ProviderHandler) DeleteProvider(c *gin.Context, id string) {
	err := h.db.DeleteProvider(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
