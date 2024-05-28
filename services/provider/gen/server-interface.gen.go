// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// NewProvider defines model for NewProvider.
type NewProvider struct {
	Bio   string `json:"bio"`
	Email string `json:"email"`

	// Firstname first and middle names of the provider
	Firstname string `json:"firstname"`

	// Image URL to the provider's profile image
	Image     *string  `json:"image,omitempty"`
	Languages []string `json:"languages"`

	// Lastname last name of the provider
	Lastname string   `json:"lastname"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Services []string `json:"services"`
	Suffix   string   `json:"suffix"`
}

// Provider defines model for Provider.
type Provider struct {
	Bio   string `json:"bio"`
	Email string `json:"email"`

	// Firstname first and middle names of the provider
	Firstname string `json:"firstname"`
	Id        string `json:"id"`

	// Image URL to the provider's profile image
	Image     *string  `json:"image,omitempty"`
	Languages []string `json:"languages"`

	// Lastname last name of the provider
	Lastname string `json:"lastname"`

	// Password This field is never returned in a response.
	Password string   `json:"-"`
	Phone    string   `json:"phone"`
	Services []string `json:"services"`
	Suffix   string   `json:"suffix"`
}

// ProviderLogin defines model for ProviderLogin.
type ProviderLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ProviderUpdate defines model for ProviderUpdate.
type ProviderUpdate struct {
	Bio   *string `json:"bio,omitempty"`
	Email *string `json:"email,omitempty"`

	// Firstname first and middle names of the provider
	Firstname *string `json:"firstname,omitempty"`

	// Image URL to the provider's profile image
	Image     *string   `json:"image,omitempty"`
	Languages *[]string `json:"languages,omitempty"`

	// Lastname last name of the provider
	Lastname *string   `json:"lastname,omitempty"`
	Password *string   `json:"password,omitempty"`
	Phone    *string   `json:"phone,omitempty"`
	Services *[]string `json:"services,omitempty"`
	Suffix   *string   `json:"suffix,omitempty"`
}

// UpdateProvider defines model for UpdateProvider.
type UpdateProvider = ProviderUpdate

// GetProvidersParams defines parameters for GetProviders.
type GetProvidersParams struct {
	// Name Filter providers by name
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Service Filter providers by service
	Service *string `form:"service,omitempty" json:"service,omitempty"`
}

// CreateProviderJSONRequestBody defines body for CreateProvider for application/json ContentType.
type CreateProviderJSONRequestBody = NewProvider

// ProviderLoginJSONRequestBody defines body for ProviderLogin for application/json ContentType.
type ProviderLoginJSONRequestBody = ProviderLogin

// UpdateProviderJSONRequestBody defines body for UpdateProvider for application/json ContentType.
type UpdateProviderJSONRequestBody = ProviderUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Health check
	// (GET /health)
	CheckHealth(c *gin.Context)
	// Get all providers
	// (GET /provider)
	GetProviders(c *gin.Context, params GetProvidersParams)
	// Create a provider
	// (POST /provider)
	CreateProvider(c *gin.Context)
	// Login as a provider
	// (POST /provider/login)
	ProviderLogin(c *gin.Context)
	// Delete a provider by ID
	// (DELETE /provider/{id})
	DeleteProvider(c *gin.Context, id string)
	// Get a provider by ID
	// (GET /provider/{id})
	GetProvider(c *gin.Context, id string)
	// Update a provider by ID
	// (PATCH /provider/{id})
	UpdateProvider(c *gin.Context, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// CheckHealth operation middleware
func (siw *ServerInterfaceWrapper) CheckHealth(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CheckHealth(c)
}

// GetProviders operation middleware
func (siw *ServerInterfaceWrapper) GetProviders(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProvidersParams

	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", c.Request.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "service" -------------

	err = runtime.BindQueryParameter("form", true, false, "service", c.Request.URL.Query(), &params.Service)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter service: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProviders(c, params)
}

// CreateProvider operation middleware
func (siw *ServerInterfaceWrapper) CreateProvider(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateProvider(c)
}

// ProviderLogin operation middleware
func (siw *ServerInterfaceWrapper) ProviderLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ProviderLogin(c)
}

// DeleteProvider operation middleware
func (siw *ServerInterfaceWrapper) DeleteProvider(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteProvider(c, id)
}

// GetProvider operation middleware
func (siw *ServerInterfaceWrapper) GetProvider(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetProvider(c, id)
}

// UpdateProvider operation middleware
func (siw *ServerInterfaceWrapper) UpdateProvider(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateProvider(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/health", wrapper.CheckHealth)
	router.GET(options.BaseURL+"/provider", wrapper.GetProviders)
	router.POST(options.BaseURL+"/provider", wrapper.CreateProvider)
	router.POST(options.BaseURL+"/provider/login", wrapper.ProviderLogin)
	router.DELETE(options.BaseURL+"/provider/:id", wrapper.DeleteProvider)
	router.GET(options.BaseURL+"/provider/:id", wrapper.GetProvider)
	router.PATCH(options.BaseURL+"/provider/:id", wrapper.UpdateProvider)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXQW/rNgz+K4I2YBenydt28u2twbpgXRcM66noQbVpm60tqZKSNij83wdJtmPHSpts",
	"wVo89BRHtEl+nz6S0gtNRCUFB240jV+ogscVaPOLSBHcwhU8LZVYYwrK/k0EN8CNfWRSlpgwg4JP77Xg",
	"dk0nBVTMPn2vIKMx/W669T/1Vj3t+6zrOnJRUUFKY6NWUEe0NV+KHPnJ4g69hiNfy5QZODnk1qF3H4xd",
	"R42vEO9SCQnKNJtyh8L+mI0EGlNtFPKc1hGFimEZtGSotOGsAmtNQScKpYVBY28ijKekwjQtgdjXNBEZ",
	"MQUQ2SYRjZ1ixfKAw+u/LokRg69/0PYxwxKI/yjgrWQ8X7HcI0QDlQ4iaRaYUmzjP9sHzFocmEOwSKb1",
	"k1BpMKYsBIegRYNaY3JsznqVZfhsX4VnVsnSGv+YR2RZzCMCJhknOJDLTW87ewR0fntp9WmNnG5aNK1a",
	"etBvu7Di7h4SQ3uV+HFFGN6yb1WbQzd/F6hJhlCmBDXhsAZFFJiV4pAS5IQRBVoKruFs5Diiz5NcTGwf",
	"m2DOhQLfiJp1+/JEP6CcCBeNlRMpkBsrhaZXfriywJRGp6yN1wqiG03Dqtiv/lc6zA6MowqzmSifM+Jz",
	"RoxUYpeQZ2KM++tyQTKhSMU4y5HnHW5i31eVO+PYEGhc5FZt5OtyQSO6BqW9oy9ns7OZzVdI4EwijelP",
	"bsnSZQqHdVoAK01hH3NwBykrVBdikdKYnheQPPzm37GV4DuW+/TH2Wyc/J+/O7B6VVVMbWhM/bcksX6c",
	"aSp7YysY8wJMi0m7XBWrwNg/8c1uvF+xNKA6ijS525Cmq9gWQB9XoDY0ol5itG043cFwtE+HBGhUsyfG",
	"1ro/zG2Yy4PPs51WDznYjiVs9+GNfbsAQ1hZboG7QhI6pBEF/XN51LupbPblOLjMDG8dI2q+nPyoH2LA",
	"o0h3aPCrhG2bz0DD07IbNUFqhhPpXzCzcy/6j7IZjiEjHoCHR16gWb2hF5cgYXo/VS+Y1r5hlOCn4pCr",
	"uVvvyejVwu+63mLe1qFtatsydAeO4T3uuIr8edzdrgQ5b7geove597DbNrGY2w17q8W9P9DZ/1Jf4Q4T",
	"IEwykxRjynZu/+/B2pG1u5Nx/VGI93kFuK/r+p8AAAD//z5DPAV1EgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
