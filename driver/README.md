# Data Store

Connect with the required database which we currently using for our cms

## Supported Databases

- MySQL
- ClickHouse
- Redis
- Aerospike

## Installation

Install go using go get

```bash
go get -u github.com/lemmamedia/data-store
```

## Usage

The "data-store" package provides a unified interface to connect with various databases, including MySQL, ClickHouse, Redis, and Aerospike. Each database connector implements the Connector interface, allowing you to perform common database operations with consistent methods.

### MySQL Example

```go
package main

import (
    "github.com/lemmamedia/data-store/mysql"
)

func main() {
    // Set up your MySQL connection parameters
    config := mysql.Config{
        Host:     "localhost",
        Port:     3306,
        Username: "root",
        Password: "password",
        Database: "mydb",
    }

    // Connect to MySQL
    conn, err := mysql.Connect(config)
    if err != nil {
        // Handle error
    }

    // Perform database operations using the "conn" object
}
```

### ClickHouse Example

```go
package main

import (
    "github.com/lemmamedia/data-store/clickhouse"
)

func main() {
    // Set up your ClickHouse connection parameters
    config := clickhouse.Config{
        Host:     "localhost",
        Port:     9000,
        Username: "default",
        Password: "",
    }

    // Connect to ClickHouse
    conn, err := clickhouse.Connect(config)
    if err != nil {
        // Handle error
    }

    // Perform database operations using the "conn" object
}
```

### Redis Example

```go
package main

import (
    "github.com/lemmamedia/data-store/redis"
)

func main() {
    // Set up your Redis connection parameters
    config := redis.Config{
        Host:     "localhost",
        Port:     6379,
        Password: "",
    }

    // Connect to Redis
    conn, err := redis.Connect(config)
    if err != nil {
        // Handle error
    }

    // Perform database operations using the "conn" object
}
```

### Aerospike Example

```go
package main

import (
    "github.com/lemmamedia/data-store/aerospike"
)

func main() {
    // Set up your Aerospike connection parameters
    config := aerospike.Config{
        Host:     "localhost",
        Port:     3000,
        Namespace: "test",
        Set:      "demo",
    }

    // Connect to Aerospike
    conn, err := aerospike.Connect(config)
    if err != nil {
        // Handle error
    }

    // Perform database operations using the "conn" object
}
```
