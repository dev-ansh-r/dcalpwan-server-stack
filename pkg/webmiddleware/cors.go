package webmiddleware

import "github.com/gorilla/handlers"

// CORSConfig is the configuration for the CORS middleware.
type CORSConfig struct {
	AllowedHeaders   []string
	AllowedMethods   []string
	AllowedOrigins   []string
	ExposedHeaders   []string
	MaxAge           int
	AllowCredentials bool
}

func (c CORSConfig) options() []handlers.CORSOption {
	var options []handlers.CORSOption
	if len(c.AllowedHeaders) > 0 {
		options = append(options, handlers.AllowedHeaders(c.AllowedHeaders))
	}
	if len(c.AllowedMethods) > 0 {
		options = append(options, handlers.AllowedMethods(c.AllowedMethods))
	}
	if len(c.AllowedOrigins) > 0 {
		options = append(options, handlers.AllowedOrigins(c.AllowedOrigins))
	}
	if len(c.ExposedHeaders) > 0 {
		options = append(options, handlers.ExposedHeaders(c.ExposedHeaders))
	}
	if c.MaxAge > 0 {
		options = append(options, handlers.MaxAge(c.MaxAge))
	}
	if c.AllowCredentials {
		options = append(options, handlers.AllowCredentials())
	}
	return options
}

// CORS returns a middleware that handles Cross-Origin Resource Sharing.
func CORS(config CORSConfig) MiddlewareFunc {
	return MiddlewareFunc(handlers.CORS(config.options()...))
}
