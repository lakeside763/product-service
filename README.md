## Product-Service


### Author
- By Moses Idowu

### Overview
Product service for managing product categories, discounts and products search.

### Prerequisites
The solution was built using the following
- Golang (1.22 Go version)
- PostgreSQL
- Redis Cache

### Installation, Deployment and Testing
Easily run on docker

Steps
- Clone the repo
```
- git clone https://github.com/lakeside763/product-service.git
- cd product-service
```

- Setup env configuration
Using env-sample.txt file to setup .env as follows
```
DATABASE_URL=postgres://postgres:password@postgres:5432/mytheresa?sslmode=disable
DB_HOST=postgres
DB_NAME=mytheresa
DB_USER=postgres
DB_PASSWORD=password
DB_PORT=5432
REDIS_URL=redis://redis:6379
APP_PORT=4500
```

- Run the app on Docker using docker compose cmd
```
- docker compose up
```

- Run a coverage test (unit testing the product repo and service)
```
go test -v ./...  
```

### API Documentation
- GetProducts endpoint
```
- http://localhost:4500/products?category=boots&priceLessThan=800&pageSize=10
From the query string
- category is required
- priceLessThan is optional
- pageSize is optional
```
- GetProduct sample results
```
[
    {
        "id": "2c870ffa-a71e-48f9-9e98-20b0f640c5b6",
        "sku": "000003",
        "name": "Ashlington leather ankle boots",
        "category": "boots",
        "price": {
            "original": 710,
            "final": 461.5,
            "discount_percentage": "35%",
            "currency": "EUR"
        },
        "created_at": "2024-11-10T06:16:09.643039Z",
        "updated_at": "2024-11-10T06:16:09.643039Z"
    },
    {
        "id": "44c3472b-978e-4e50-8acf-676d171a2721",
        "sku": "000004",
        "name": "Naima embellished suede sandals",
        "category": "sandals",
        "price": {
            "original": 795,
            "final": 636,
            "discount_percentage": "20%",
            "currency": "EUR"
        },
        "created_at": "2024-11-10T06:16:09.643039Z",
        "updated_at": "2024-11-10T06:16:09.643039Z"
    },
    {
        "id": "e15e3796-d947-4497-9e13-6a850da634ce",
        "sku": "000005",
        "name": "Nathane leather sneakers",
        "category": "sneakers",
        "price": {
            "original": 590,
            "final": 442.5,
            "discount_percentage": "25%",
            "currency": "EUR"
        },
        "created_at": "2024-11-10T06:16:09.643039Z",
        "updated_at": "2024-11-10T06:16:09.643039Z"
    }
]
```

### Folder structure
- Using Hexagonal Architecture to manage the product-service experience
```
/cmd
  /app
    - main.go             # Entry point for the application


/config                   # Application configuration
  config.go

/init-scripts             # Postgres db configuration, table and sample data insertion
  init.sql


/internal
  /adapters                         # Managing external dependencies
    /cache
      redis.go
    /database
      /migrations
      postgres_conn.go
    /repositories
      data_repo.go
      product_repo.go
      product_repo_test.go

  /core                           # Managing app business logic
    /models
      product_model.go
    /services
      product_service.go
      product_service_test.go

  /ports                           # Managing app input dependencies
    /http
      /handlers
        product_handler.go
      /routes
        product_route.go
    /interfaces
      /mocks
      product_interface.go
      redis_interface.go

/pkg
  /utils                        # Managing helper functions
    json_response.go
    price_conversion.go

env-sample.txt                  # env data samples
```

