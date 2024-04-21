package main

import (
	"fmt"
	"os"
)

// Configurer defines an interface that all concrete configs and decorators will implement
// this interface will allow us to embed decorators in other decorators and reload the configuration
type Configurer interface {
	Reload() error
}

/*
#########################################################################
# Base Config Section
#########################################################################
*/

// Config is the base configuration struct for our application
type Config struct {
	Address string
	Port    string
}

// Reload reloads the configuration from the environment variables and implements the Configurer interface
func (c *Config) Reload() error {
	fmt.Println("Reloading base config")

	// Load the environment variables
	c.Address = os.Getenv("ADDRESS")
	c.Port = os.Getenv("PORT")

	// Check if the environment variables are set and not empty strings
	if c.Address == "" {
		// set a default address if the environment variable is not set
		c.Address = "http://localhost"
	}
	if c.Port == "" {
		// set a default port if the environment variable is not set
		c.Port = "8081"
	}
	return nil
}

// NewConfig creates a new Config struct
func NewConfig(address string, port string) *Config {
	return &Config{
		Address: address,
		Port:    port,
	}
}

/*
#########################################################################
# Database Config Section - Decorator for the Config struct
#########################################################################
*/

// DatabaseConfig is a decorator for the Config struct
type DatabaseConfig struct {
	Configurer
	DBAddress string
	DBPort    string
}

// NewDatabaseConfig creates a new DatabaseConfig struct dependecy inject the Configurer interface
// this will allow us to reload the base configuration when the DatabaseConfig is reloaded
func NewDatabaseConfig(config Configurer, dbAddress string, dbPort string) *DatabaseConfig {
	return &DatabaseConfig{
		Configurer: config,
		DBAddress:  dbAddress,
		DBPort:     dbPort,
	}
}

// Reload reloads the configuration from the environment variables and implements the Configurer interface
func (d *DatabaseConfig) Reload() error {
	fmt.Println("Reloading database config")

	// Reload the base configuration
	if err := d.Configurer.Reload(); err != nil {
		return err
	}

	// Load the environment variables
	d.DBAddress = os.Getenv("DB_ADDRESS")
	d.DBPort = os.Getenv("DB_PORT")

	// Check if the environment variables are set and not empty strings
	if d.DBAddress == "" {
		// set a default address if the environment variable is not set
		d.DBAddress = "http://localhost"
	}
	if d.DBPort == "" {
		// set a default port if the environment variable is not set
		d.DBPort = "37017"
	}

	return nil
}

/*
#########################################################################
# Message of the Day Config Section - Decorator for the Config struct
#########################################################################
*/

// MessageOfTheDay is a decorator for the Config struct and adds a message of the day
// functionality to the configuration
type MessageOfTheDay struct {
	Configurer
	MOTD string
}

// NewMessageOfTheDay creates a new MessageOfTheDay struct that decorates the Config struct
func NewMessageOfTheDay(config Configurer, motd string) *MessageOfTheDay {
	return &MessageOfTheDay{
		Configurer: config,
		MOTD:       motd,
	}
}

// Reload reloads the configuration from the environment variables and implements the Configurer interface
func (m *MessageOfTheDay) Reload() error {
	fmt.Println("Reloading message of the day")

	// Reload the base configuration
	if err := m.Configurer.Reload(); err != nil {
		return err
	}

	// Load the environment variables
	m.MOTD = os.Getenv("MOTD")

	if m.MOTD == "" {
		// set a default message of the day if the environment variable is not set
		m.MOTD = "Have a Nice Day!"
	}

	return nil
}

// main function which is where the application will begin execution

func main() {
	// Create a new Config and DatabaseConfig note the current values before reloading
	config := NewConfig("http://webapp", "8080")
	dbConfig := NewDatabaseConfig(config, "http://mongodb", "27017")
	motdConfig := NewMessageOfTheDay(dbConfig, "Hello, World!")

	// Print the current values
	fmt.Printf("Config Address: %s, Port: %s\n", config.Address, config.Port)
	fmt.Printf("Database Address: %s, Port: %s\n", dbConfig.DBAddress, dbConfig.DBPort)
	fmt.Printf("Message of the Day: %s\n", motdConfig.MOTD)

	// Reload the last decorator in the chain which will reload
	// all the decorators.
	if err := motdConfig.Reload(); err != nil {
		fmt.Printf("Error reloading configuration: %v\n", err)
		return
	}

	// Print the new values reloaded from the environment variables
	fmt.Printf("Web App Address: %s, Port: %s\n", config.Address, config.Port)
	fmt.Printf("Database Address: %s, Port: %s\n", dbConfig.DBAddress, dbConfig.DBPort)
	fmt.Printf("Message of the Day: %s\n", motdConfig.MOTD)

	/*
		Additional Notes
		The Config struct can be used as a standalone configuration struct
		The DatabaseConfig struct can be used as a standalone configuration struct
		The MessageOfTheDay struct can be used as a standalone configuration struct
		The DatabaseConfig and MessageOfTheDay structs can be combined to create a configuration with both database and message of the day functionality
		You could use the final struct that implements the Configurer interface to use in the context of an application.

		Example:

		type Service struct {
			Configurer
			*http.Server
		}

		You could then just call Service.Reload() to reload the configuration for the service thus reloading all the decorators in the chain.
	*/

}
