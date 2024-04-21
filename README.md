# ConfigDecorator

This example demonstrates the Decorator pattern. This is a structural pattern
used to add new behaviors to objects dynamically by placing them inside special
wrapper objects containing these behaviors. This pattern creates a flexible design
that is easy to extend without modifying existing code.

In this example:
- The **Config** struct holds the base configuration for our application.
- The **DatabaseConfig** and **MessageOfTheDay** structs act as decorators to the 'Config' struct,
  adding database configuration and a message of the day functionality, respectively.
- All configuration structs implement the **Configurer** interface which includes a 'Reload' method
  for reloading configuration from the environment variables.

This allows the **DatabaseConfig** and **MessageOfTheDay** decorators to reuse and extend
the **Reload** method of the **Config** struct dynamically, demonstrating the Decorator pattern's flexibility.


**TLDR:** The Decorator pattern allows you to add new behaviors to objects dynamically by embedding
them inside other types or structs. By defining a common interface for all decorators, you can
easily extend the functionality of an object without modifying its core implementation.

## Prerequisites
- Go 1.22.2 or later

## Getting Started

To run this example, simply execute the following command:

```bash
go run main.go
```

### Output
```bash
go run main.go 
Config Address: http://webapp, Port: 8080
Database Address: http://mongodb, Port: 27017
Message of the Day: Hello, World!
Reloading message of the day
Reloading database config
Reloading base config
Web App Address: http://localhost, Port: 8081
Database Address: http://localhost, Port: 37017
Message of the Day: Have a Nice Day!
```
