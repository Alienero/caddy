package server

import (
	"net"

	"github.com/mholt/caddy/middleware"
)

// Config configuration for a single server.
type Config struct {
	// The hostname or IP on which to serve
	Host string

	// The host address to bind on - defaults to (virtual) Host if empty
	BindHost string `json:",omitempty"`

	// The port to listen on
	Port string

	// The directory from which to serve files
	Root string

	// HTTPS configuration
	TLS TLSConfig

	// Middleware stack; map of path scope to middleware -- TODO: Support path scope?
	Middleware map[string][]*middleware.Middleware `json:"-"`

	// Map of middleware to directive name -- TODO: Support path scope?
	MiddlewareMap map[*middleware.Middleware]string `json:"-"`

	// Map of directive name to pointer to handler -- TODO: Support path scope?
	HandlerMap map[string]middleware.Handler

	// Functions (or methods) to execute at server start; these
	// are executed before any parts of the server are configured,
	// and the functions are blocking
	Startup []func() error `json:"-"`

	// Functions (or methods) to execute when the server quits;
	// these are executed in response to SIGINT and are blocking
	Shutdown []func() error `json:"-"`

	// The path to the configuration file from which this was loaded
	ConfigFile string

	// The name of the application
	AppName string

	// The application's version
	AppVersion string
}

// Address returns the host:port of c as a string.
func (c Config) Address() string {
	if c.BindHost != "" {
		return net.JoinHostPort(c.BindHost, c.Port)
	}
	return net.JoinHostPort(c.Host, c.Port)
}

// TLSConfig describes how TLS should be configured and used,
// if at all. A certificate and key are both required.
// The rest is optional.
type TLSConfig struct {
	Enabled                  bool     `json:",omitempty"`
	Certificate              string   `json:",omitempty"`
	Key                      string   `json:",omitempty"`
	Ciphers                  []uint16 `json:",omitempty"`
	ProtocolMinVersion       uint16   `json:",omitempty"`
	ProtocolMaxVersion       uint16   `json:",omitempty"`
	PreferServerCipherSuites bool     `json:",omitempty"`
	ClientCerts              []string `json:",omitempty"`
}
