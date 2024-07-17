# Sample GO REST API

## Setup environment

- Install golang 1.22.x: <https://go.dev/dl/>
- Install wire `go install github.com/google/wire/cmd/wire@latest`

## Building

```bash
go build
./gin-gonic-api
```

### Update dependency wiring (if needed)

```bash
go generate ./...
```

## Testing

```bash
go test ./... -coverprofile=coverage.out
```

## Accessing the app

Either via browser on <http://localhost:8080> or via `curl`:

```bash
curl -X GET http://localhost:8080/api/accounts
curl -X POST --header "Content-Type: application/json" --data '{"name":"Test Account","email":"test@foo.com"}' http://localhost:8080/api/accounts
curl -X GET http://localhost:8080/api/accounts/1
```

## Quickstart Gin Web Framework:

https://gin-gonic.com/docs/quickstart/

## Very simple Compile Daemon for Go:

Watches your .go files in a directory and invokes go build if a file changed. Nothing more.

https://github.com/githubnemo/CompileDaemon

## GoDotEnv

Loads env vars from a .env file

https://github.com/joho/godotenv

## VSCode extensions:

Go

## How to start your first project

### 1. Create Workspace

Create a directory for your Go module source code.

```console
$ mkdir go-crud
$ cd go crud
```

Start your module using the go mod init command.

```console
$ go mod init

```

The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.

### Compile Deamon for Go:

Watches your .go files in a directory and invokes go build if a file changed. Nothing more.

#### Installation

```console
$ go get github.com/githubnemo/CompileDaemon
$ go install github.com/githubnemo/CompileDaemon
```

## GoDotEnv

A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).

### Installation

```console
go get github.com/joho/godotenv
```

## Gin Web Framework

Gin, a HTTP web framework written in Golang, provides features like routing and middleware out of the box. Gin helps developers to reduce boilerplate code and improve productivity - simplifying the process of building microservices.

```console
$ go get -u github.com/gin-gonic/gin
```

## GORM library for Golang

GORM simplifies database interactions by mapping Go structs to database tables.

go get -u gorm.io/gorm

```console
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/postgres
```

# main.go

```console
$ mkdir main.go
```

```go
// main.go
package main

import (
	"example.com/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test123d‚",
		})
	})
	r.Run()
}
```

In its simplest form, the defaults will do. With the current working directory set to the source directory you can simply…

```console
$ CompileDaemon
```

… and it will recompile your code whenever you save a source file.

If you want it to also run your program each time it builds you might add…

```console
$ CompileDaemon -command="./go-crud"
```

… and it will also keep a copy of your program running. Killing the old one and starting a new one each time you build.

## Open Browser

http://localhost:2000/

## Dot-Env-File

Dotenv is a zero-dependency module that loads environment variables in a .env file. Storing configuration in the environment separate from code is based on The Twelve-Factor App methodology.

Set PORT in .env file:

```.env
PORT=3000
```

Load variables from .env

```terminal
$ mkdir initializers
$ touch loadEnvVariables.go
```

```go
// loadEnvVariables.go
package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

```

initialize Variables

```go

// main.go
package main

import (
    "example.com/go-crud/initializers"
	)
    "github.com/gin-gonic/gin"

// executed before the main function
func init() {
    initializers.LoadEnvVariables()
}

func main() {
	// Creates a default Gin engine with logger and recovery middleware
    r := gin.Default()

    // Defines a route for GET requests on the root path
    r.GET("/", func(c *gin.Context) {
        // Sends a JSON response with the key "message" and value "Test123d"
        c.JSON(200, gin.H{
            "message": "Test123d‚",
        })
    })

    // Starts the server, listening on 0.0.0.0:2000 (configured in .env)
    r.Run()
}
```

## Connect to database

create new file in initializers folder

```console
$ touch database.go
```

```go
// database.go

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5435 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	} else {
		log.Fatal("Connected to database!")
	}
}
```

create dsn variable in .env file

```go
PORT=2000
DB_URL="host=localhost user=postgres password=postgres dbname=postgres port=5435 sslmode=disable"
```

```go
// add DB_URL variable
func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	} else {
		log.Fatal("Connected to database!")
	}
}

```

create models folder

```console
mkdir models
touch postModel.go
```

```go
// postModel.go
package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body string
}
```

create migrate folder

```console
$ mkdir migrate
$ touch migrate.go
```

```go
// migrate.go
package main

import (
	"example.com/go-crdud/initializers"
	"example.com/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
```

```console
go run migrate/migrate.go
```

## PostgreSQL

we will be using PostgreSQL as our database. So, for development purposes, you will need to install it in your local environment. However, in production, you might consider a more robust and secure solution such as a cloud offering. An example of this is AWS Aurora. You can download PostgreSQL from the official website here:

https://www.postgresql.org/download/

## pgAdmin 4

this is a user interface that allows us to manage our Postgres database visually. You can download pgAdmin here:

https://www.pgadmin.org/download/

Weiter gehts bei Minute 22
https://www.youtube.com/watch?v=lf_kiH_NPvM&t=251s
