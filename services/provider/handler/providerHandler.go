package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/services/provider/gen"
	"github.com/services/provider/security"
)

// ProviderHandler defines the handlers for the Provider service
type ProviderHandler struct {
	db ProviderDatabase
}

// ProviderDatabase defines the database operations required for the Provider service
type ProviderDatabase interface {
	GetPassword(email string) (string, error)
	GetProviderID(email string) (string, error)
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

// ProviderLogin authenticates a provider and returns a JWT. It expects a gen.ProviderLogin in the request
func (h *ProviderHandler) ProviderLogin(c *gin.Context) {
	var providerLogin gen.ProviderLogin
	if err := c.BindJSON(&providerLogin); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Retrieve hashed password
	storedPass, err := h.db.GetPassword(providerLogin.Email)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// Compare hashed password with provider password
	err = security.ComparePasswords(storedPass, providerLogin.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// Retrieve provider ID
	providerId, err := h.db.GetProviderID(providerLogin.Email)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT
	token, err := security.GenerateJWT(providerId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
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

	// Hash password
	hashedPass, err := security.HashPassword(provider.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	provider.Password = hashedPass

	// Create provider
	createdProvider, err := h.db.CreateProvider(&provider)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, createdProvider)
}

// DeleteProvider calls ProviderDatabase.DeleteProvider() and returns the result. It expects an id parameter
func (h *ProviderHandler) DeleteProvider(c *gin.Context, id string) {
	security.AuthMiddleware()(c)
	if c.IsAborted() {
		return
	}

	err := h.db.DeleteProvider(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// TODO: write UpdateProvider in open api yaml and in the handler/db
