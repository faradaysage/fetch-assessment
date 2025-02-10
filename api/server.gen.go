//go:build go1.22

// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Item defines model for Item.
type Item struct {
	// Price The total price payed for this item.
	Price string `json:"price"`

	// ShortDescription The Short Product Description for the item.
	ShortDescription string `json:"shortDescription"`
}

// Receipt defines model for Receipt.
type Receipt struct {
	Items []Item `json:"items"`

	// PurchaseDate The date of the purchase printed on the receipt.
	PurchaseDate openapi_types.Date `json:"purchaseDate"`

	// PurchaseTime The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime"`

	// Retailer The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer"`

	// Total The total amount paid on the receipt.
	Total string `json:"total"`
}

// PostReceiptsProcessJSONRequestBody defines body for PostReceiptsProcess for application/json ContentType.
type PostReceiptsProcessJSONRequestBody = Receipt

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Submits a receipt for processing.
	// (POST /receipts/process)
	PostReceiptsProcess(w http.ResponseWriter, r *http.Request)
	// Returns the points awarded for the receipt.
	// (GET /receipts/{id}/points)
	GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostReceiptsProcess operation middleware
func (siw *ServerInterfaceWrapper) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostReceiptsProcess(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetReceiptsIdPoints operation middleware
func (siw *ServerInterfaceWrapper) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", r.PathValue("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReceiptsIdPoints(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/receipts/process", wrapper.PostReceiptsProcess)
	m.HandleFunc("GET "+options.BaseURL+"/receipts/{id}/points", wrapper.GetReceiptsIdPoints)

	return m
}

type BadRequestResponse struct {
}

type NotFoundResponse struct {
}

type PostReceiptsProcessRequestObject struct {
	Body *PostReceiptsProcessJSONRequestBody
}

type PostReceiptsProcessResponseObject interface {
	VisitPostReceiptsProcessResponse(w http.ResponseWriter) error
}

type PostReceiptsProcess200JSONResponse struct {
	Id string `json:"id"`
}

func (response PostReceiptsProcess200JSONResponse) VisitPostReceiptsProcessResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostReceiptsProcess400Response = BadRequestResponse

func (response PostReceiptsProcess400Response) VisitPostReceiptsProcessResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type GetReceiptsIdPointsRequestObject struct {
	Id string `json:"id"`
}

type GetReceiptsIdPointsResponseObject interface {
	VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error
}

type GetReceiptsIdPoints200JSONResponse struct {
	Points *int64 `json:"points,omitempty"`
}

func (response GetReceiptsIdPoints200JSONResponse) VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetReceiptsIdPoints404Response = NotFoundResponse

func (response GetReceiptsIdPoints404Response) VisitGetReceiptsIdPointsResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Submits a receipt for processing.
	// (POST /receipts/process)
	PostReceiptsProcess(ctx context.Context, request PostReceiptsProcessRequestObject) (PostReceiptsProcessResponseObject, error)
	// Returns the points awarded for the receipt.
	// (GET /receipts/{id}/points)
	GetReceiptsIdPoints(ctx context.Context, request GetReceiptsIdPointsRequestObject) (GetReceiptsIdPointsResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostReceiptsProcess operation middleware
func (sh *strictHandler) PostReceiptsProcess(w http.ResponseWriter, r *http.Request) {
	var request PostReceiptsProcessRequestObject

	var body PostReceiptsProcessJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostReceiptsProcess(ctx, request.(PostReceiptsProcessRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostReceiptsProcess")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostReceiptsProcessResponseObject); ok {
		if err := validResponse.VisitPostReceiptsProcessResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetReceiptsIdPoints operation middleware
func (sh *strictHandler) GetReceiptsIdPoints(w http.ResponseWriter, r *http.Request, id string) {
	var request GetReceiptsIdPointsRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetReceiptsIdPoints(ctx, request.(GetReceiptsIdPointsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetReceiptsIdPoints")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetReceiptsIdPointsResponseObject); ok {
		if err := validResponse.VisitGetReceiptsIdPointsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xWQW/bOBP9K8R8vVW2ZcWf0ei2XWMXxiJFkPQWZQFaHMXsRiQ7HDU1Av33BSVZtmyl",
	"SYA9JZKGM4/vvZnxM+S2dNagYQ/pMxB6Z43H5uGzVDf4vULP4Umhz0k71tZACl+3KAhz1I6F9kKbH/JR",
	"qynUEXyx/IetjDo/9MX2Z4oQIQpLgreSxXo1hbqOwOdbLGVTfc1Yhr+OrENi3WJypHMch8OW5aNoAoST",
	"O9ynD/AYyylEgD9l6R4RUlhOF5cQgZPMSCHD31mmPmbZNMvUc1J/gAh450KkZ9LmIVzMby3x6rjuGIzb",
	"ECWuyaoqZ3EU3sHBETRXtjIstRErfBLz5PqvIbS7LHvKMp9lk/uPI8jqCAi/V5pQQXp3DjPqWLvvT9rN",
	"N8w53Omm1eOc6ABy+M8HwgJS+N/sYJlZp9esEauOoNRm3cbP+2KSSO7CR1dRvpUeV5JfkFBJRmGLhqV9",
	"dFDUMCphTfO+c9CQwCROkkk8n8RziKCwVEqGFEK6MSH3qb/q8iUv6fLNQESymGxtRe0h/OkwZ1RDfPOL",
	"dAgtxI5BI2SpH5HGYRl5gLWPFJaEZ0t4DCr0ZEH21GZZFcfJ8kr8bskgiStJ/yC/5LU2eNRxETTN9qs+",
	"lGXwtHBS/1q59zfiid17xk4MdiJz1Bl5D/28GUJibQp7fqvfhNcBbs+uI5uj9zYUZc3NRbpOCp3ff/uB",
	"5NsU82k8jQNx1qGRTkMKF9N4etFefds02KxL72dd/qYr7dj4va02pWYv5NFApT0sbR4Cx6GbZYhfK0jh",
	"2nruIPoOIrREoufPVu1CkdwaRtPUk8496rw5P/vm22HXdvtrs2A/UuqhUkwVNi+ONkwSx+8qezKimhVz",
	"sJJUm+Xm/8t4EiMWk0WyySeXar6cqGLxqbiI8dPlJjm12u0bBqpWL7hlqMkNckXGN1Zfr4T0Xj8YVILt",
	"0P11BIv23mMs9vzMjtZvsxqrspS0e5P4If7gpmet6pmzutvxDzhiqGPwbaiQT5JUv0UH7Tu01p/YO2ut",
	"rts6gWaSJTKSh/RubE6sV4dR1mfW4WvoCIggTDtIA/2nPoqOPfG6nPf/qe0OTPbWm8fx0WzXhpeLA4yw",
	"NB6QGlFeNVEz5atyE+Z6caJEZ53F69bpf4INjfMekeu6rv8NAAD//+zUmakeCgAA",
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
