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

// Defines values for Gender.
const (
	Female Gender = "female"
	Male   Gender = "male"
)

// Allergy defines model for Allergy.
type Allergy struct {
	// Description Description of the allergy
	Description *string `json:"description,omitempty"`

	// Name Name of the allergy
	Name string `json:"name"`
}

// Gender defines model for Gender.
type Gender string

// JWT defines model for JWT.
type JWT struct {
	// Id ID of the patient
	Id    string `json:"id"`
	Token string `json:"token"`
}

// NewPatient defines model for NewPatient.
type NewPatient struct {
	Allergies *[]Allergy `json:"allergies,omitempty"`
	Birth     string     `json:"birth"`
	Email     string     `json:"email"`

	// Firstname First and middle names of the patient
	Firstname string `json:"firstname"`
	Gender    Gender `json:"gender"`

	// Language Preferred language of the patient
	Language *string `json:"language,omitempty"`

	// Lastname Last name of the patient
	Lastname      string          `json:"lastname"`
	Password      string          `json:"password"`
	Phone         *string         `json:"phone,omitempty"`
	Prescriptions *[]Prescription `json:"prescriptions,omitempty"`
}

// Patient defines model for Patient.
type Patient struct {
	Allergies *[]Allergy `json:"allergies,omitempty"`
	Birth     string     `json:"birth"`
	Email     string     `json:"email"`

	// Firstname First and middle names of the patient
	Firstname string `json:"firstname"`
	Gender    Gender `json:"gender"`
	Id        string `json:"id"`

	// Language Preferred language of the patient
	Language *string `json:"language,omitempty"`

	// Lastname Last name of the patient
	Lastname string `json:"lastname"`

	// Password This field is never returned in a response.
	Password      string          `json:"-"`
	Phone         string          `json:"phone"`
	Prescriptions *[]Prescription `json:"prescriptions,omitempty"`
}

// PatientUpdate defines model for PatientUpdate.
type PatientUpdate struct {
	Allergies     *[]Allergy      `json:"allergies,omitempty"`
	Birth         *string         `json:"birth,omitempty"`
	Email         *string         `json:"email,omitempty"`
	Firstname     *string         `json:"firstname,omitempty"`
	Gender        *Gender         `json:"gender,omitempty"`
	Language      *string         `json:"language,omitempty"`
	Lastname      *string         `json:"lastname,omitempty"`
	Password      *string         `json:"password,omitempty"`
	Phone         *string         `json:"phone,omitempty"`
	Prescriptions *[]Prescription `json:"prescriptions,omitempty"`
}

// Prescription defines model for Prescription.
type Prescription struct {
	// Dosage Dosage of the medication
	Dosage string `json:"dosage"`

	// End Date the prescription was ended
	End *string `json:"end,omitempty"`

	// Frequency Frequency of the medication
	Frequency string `json:"frequency"`

	// Name Name of the medication
	Name string `json:"name"`

	// ProviderId ID of the provider who prescribed the medication
	ProviderId *string `db:"provider_id" json:"providerId,omitempty"`

	// Start Date the prescription was started
	Start string `json:"start"`
}

// PatientLogin defines model for PatientLogin.
type PatientLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdatePatient defines model for UpdatePatient.
type UpdatePatient = PatientUpdate

// PatientLoginJSONBody defines parameters for PatientLogin.
type PatientLoginJSONBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreatePatientJSONRequestBody defines body for CreatePatient for application/json ContentType.
type CreatePatientJSONRequestBody = NewPatient

// PatientLoginJSONRequestBody defines body for PatientLogin for application/json ContentType.
type PatientLoginJSONRequestBody PatientLoginJSONBody

// UpdatePatientJSONRequestBody defines body for UpdatePatient for application/json ContentType.
type UpdatePatientJSONRequestBody = PatientUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Health check
	// (GET /health)
	CheckHealth(c *gin.Context)
	// Get all patients
	// (GET /patient)
	GetPatients(c *gin.Context)
	// Create a patient
	// (POST /patient)
	CreatePatient(c *gin.Context)
	// Login as a patient
	// (POST /patient/login)
	PatientLogin(c *gin.Context)
	// Delete a patient by ID
	// (DELETE /patient/{id})
	DeletePatient(c *gin.Context, id string)
	// Get a patient by ID
	// (GET /patient/{id})
	GetPatient(c *gin.Context, id string)
	// Update a patient by ID
	// (PATCH /patient/{id})
	UpdatePatient(c *gin.Context, id string)
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

// GetPatients operation middleware
func (siw *ServerInterfaceWrapper) GetPatients(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPatients(c)
}

// CreatePatient operation middleware
func (siw *ServerInterfaceWrapper) CreatePatient(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreatePatient(c)
}

// PatientLogin operation middleware
func (siw *ServerInterfaceWrapper) PatientLogin(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatientLogin(c)
}

// DeletePatient operation middleware
func (siw *ServerInterfaceWrapper) DeletePatient(c *gin.Context) {

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

	siw.Handler.DeletePatient(c, id)
}

// GetPatient operation middleware
func (siw *ServerInterfaceWrapper) GetPatient(c *gin.Context) {

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

	siw.Handler.GetPatient(c, id)
}

// UpdatePatient operation middleware
func (siw *ServerInterfaceWrapper) UpdatePatient(c *gin.Context) {

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

	siw.Handler.UpdatePatient(c, id)
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
	router.GET(options.BaseURL+"/patient", wrapper.GetPatients)
	router.POST(options.BaseURL+"/patient", wrapper.CreatePatient)
	router.POST(options.BaseURL+"/patient/login", wrapper.PatientLogin)
	router.DELETE(options.BaseURL+"/patient/:id", wrapper.DeletePatient)
	router.GET(options.BaseURL+"/patient/:id", wrapper.GetPatient)
	router.PATCH(options.BaseURL+"/patient/:id", wrapper.UpdatePatient)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xY227jNhN+FYL/fynFcro367vtGk3TLra+SFGgi6CgxZHEjUSyJJ3ECPTuBUnJkkz6",
	"kDRotkCNIJB4mPnmmwOHesK5aKTgwI3Giyes4M8NaPO9oAzcwGd4WBHDgBv7lgtuukciZc1yYpjgs69a",
	"cDum8woaYp/+r6DAC/y/2SB+5mf1bCSybdvE6WQKKF4YtYE2wd3sJ1Ey/iytUgkJynTQoSGstg9mKwEv",
	"sDaK8RK3CZZE6wehaGRygudLJ2O04zbpd4j1V8gPmfCrpMTAazPXyfPCo5rbpBPlGPhQ16DKbUgNBZ0r",
	"Ji2E4BUvhzckCmQqQKQTlIRcctJAKOMzaeDk5j2unaSQ3wRfAaegnEv5prFLG1IDTnAB7uE2Auun325C",
	"sxkNkV4ve5yy81ZEmhF3wE9HC6O4XxuzY5pLU2iepB6ngUafCobeue1OE1GKuPc1U6ZyfD2SRtZ2bv7+",
	"fZZm8zSb32TZwv39HrP0cNIUTGkT9/YPdgoRTlHDKK0B2WX6DF7LnWePWdr5v01wTXi5IWUEwkpBAUoB",
	"Rf2aM9TX5JBFn4g2zoozpBypJgmWleAQn1GDxvO9vhrtCl2/F4+Dy0a2Jrui1rHfx8vRKreryv+F7stC",
	"l8Xj45uP6KmUm4ppVDCoKWIacbgHhRSYjeJAEeOIIAVaCq7hIhCc4Me0FKk98lJWcqHAH1rduF2c6jsm",
	"U+G0kTqVgnFjSe5O1TdNJlfcT2SUBxhk1pF06s7yf3tSvWplPxrV33j5Df2spr3WXiMmdDT5l268T9UG",
	"aNcwRl3EI4m6JAZ8mo/0oweikaXchvIQC5fZ5VmxULjbAc+3kTLaT50H+XTTeHy/VOKeUVDXxxu6bhV6",
	"qETPwxpoKH+gIut+aeRf/4sWNkEkS3NBoQSewqNRJDWk9D5e48UO8R+MujjRhijzHLe5DS9yXKzNTvrI",
	"G3u1RxW/4zBeiAMHAtOIIM0sKvRhdY0KoVBDOCkZL/tzRltgzDjcXe2za3GC70FpL2x+kV1k1r9CAieS",
	"4QX+zg3ZLDeVY3NWAal9NSvBMWjzyXnSRgP+WEF+96NfYw33p5HbeplloQG//Oz9sWkaorZ4gf1elFs5",
	"bmomh84nqvIKzGowMqby7NvfeUWovz0H9adNTll3BcZeyAav2GwSOkakgtEVNhl9GtgeAjj5ejC55we0",
	"zF/7Uhyz3ptA9yjwo4jsOqCxk2d1/9Uhzsrk28QLSJnsb/9mtByjxV6AzwkIhwQRfYiPJ0ZbnzU1+EZl",
	"SsjSjQ9hIokiDRhQGi++BD1tl/jXS2wLiq2LxOWpPxB8jzX9pJGM7N0vbLcBf+8ip4pAHztCp5Z75IPd",
	"aL21wFzzcjTL39zK7J9IHhspCX4Xo7Q3kAuDCrHhNFZkQlolMXkVEjv9VvYG3D4zhad427fzzoRzjyqk",
	"vW3bvwIAAP//uB8XuuUVAAA=",
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
