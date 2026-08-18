package universalsdk

import (
	"net/http"
	"strings"
	"time"
)

// CompressionType represents the compression algorithm to use
type CompressionType string

const (
	CompressionZstd CompressionType = "zstd"
	CompressionGzip CompressionType = "gzip"
)

// DefaultBaseUrl is the default API base url used when none is configured.
const DefaultBaseUrl = "https://sold-out.dev"

type Session struct {
	ApiKey string
	// BaseUrl is the API base url. If empty, DefaultBaseUrl is used.
	BaseUrl     string
	Client      *http.Client
	Compression CompressionType
}

// Default optimized HTTP client for concurrent requests
var defaultClient = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: 100,
		MaxConnsPerHost:     100,
		IdleConnTimeout:     30 * time.Second,
	},
}

// NewSession creates a new Session that can be used to make requests to the API.
func NewSession(apiKey string) *Session {
	return &Session{
		ApiKey:      apiKey,
		BaseUrl:     DefaultBaseUrl,
		Client:      defaultClient,
		Compression: CompressionZstd,
	}
}

// WithBaseUrl sets a custom API base url (e.g. "https://akamai.example.com").
// Trailing slashes are stripped. Pass an empty string to fall back to DefaultBaseUrl.
func (s *Session) WithBaseUrl(baseUrl string) *Session {
	s.BaseUrl = baseUrl
	return s
}

// baseUrl returns the configured base url, or DefaultBaseUrl if unset.
func (s *Session) baseUrl() string {
	if s.BaseUrl == "" {
		return DefaultBaseUrl
	}
	return strings.TrimRight(s.BaseUrl, "/")
}

// WithJwtKey is a no-op kept for drop-in compatibility with the upstream SDK.
// This API authenticates with the API key alone; request signing is not used.
//
// Deprecated: request signing was removed, this method does nothing.
func (s *Session) WithJwtKey(string) *Session {
	return s
}

// WithOrganization is a no-op kept for drop-in compatibility with the upstream SDK.
// This API authenticates with the API key alone; organization credentials are not used.
//
// Deprecated: organization credentials were removed, this method does nothing.
func (s *Session) WithOrganization(key, secret string) *Session {
	return s
}

// WithClient sets a new client that will be used to make requests to the API.
func (s *Session) WithClient(client *http.Client) *Session {
	s.Client = client
	return s
}

// WithCompression sets the compression type for requests.
func (s *Session) WithCompression(compression CompressionType) *Session {
	s.Compression = compression
	return s
}
