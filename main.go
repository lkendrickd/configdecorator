package main

import "fmt"

// Configurer defines an interface that all concrete configs and decorators will implement
type Configurer interface {
	GetHost() string
	GetPort() int
}

// Config is the base configuration struct for our application
type Config struct {
	address string
	port    int
}

// GetHost returns the host address
func (c *Config) GetHost() string {
	return c.address
}

// GetPort returns the port number
func (c *Config) GetPort() int {
	return c.port
}

// NewConfig creates a new Config instance
func NewConfig(address string, port int) *Config {
	return &Config{
		address: address,
		port:    port,
	}
}

// DatabaseConfig is a decorator that adds database configuration to the base Config struct
type DatabaseConfig struct {
	config    Configurer
	dbAddress string
	dbPort    int
}

// GetHost returns the host address
func (d *DatabaseConfig) GetHost() string {
	return d.config.GetHost()
}

// GetPort returns the port number
func (d *DatabaseConfig) GetPort() int {
	return d.config.GetPort()
}

// GetDBAddress returns the database address
func (d *DatabaseConfig) GetDBAddress() string {
	return d.dbAddress
}

// GetDBPort returns the port number the database is listening on
func (d *DatabaseConfig) GetDBPort() int {
	return d.dbPort
}

// NewDatabaseConfig creates a new DatabaseConfig instance
func NewDatabaseConfig(config Configurer, dbAddress string, dbPort int) *DatabaseConfig {
	return &DatabaseConfig{
		config:    config,
		dbAddress: dbAddress,
		dbPort:    dbPort,
	}
}

func main() {
	// Create a new Config instance
	config := NewConfig("http://webapp", 8080)

	// Create a new DatabaseConfig instance that decorates the Config instance with database configuration
	dbConfig := NewDatabaseConfig(config, "http://mongdb", 27017)

	// Print the host and port
	fmt.Println("Host:", dbConfig.config.GetHost())
	fmt.Println("Port:", dbConfig.config.GetPort())

	// Print the database address and port
	fmt.Println("DB Host:", dbConfig.GetDBAddress())
	fmt.Println("DB Port:", dbConfig.GetDBPort())
}
