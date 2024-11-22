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

Additionally
- docker compose build --no-cache
- docker compose down --volumes
```


- Run a coverage test (unit testing the product repo and service)
```
go test -v ./...  
```

### API Documentation
- GetProducts endpoint
```
- http://localhost:4500/products?category=boots&priceLessThan=800&pageSize=10&cursorId=MS00NjYzNTU=

From the query string
- category is required
- priceLessThan is optional
- pageSize is optional
- cursorId is optional
```
- GetProducts sample results
```
{
    "data": [
        {
            "id": "664beb14-a287-4fa5-85c2-7fbd6a5491bf",
            "sku": "000001",
            "name": "BV Lean leather ankle boots",
            "category": "boots",
            "price": {
                "original": 890,
                "final": 489.5,
                "discount_percentage": "45%",
                "currency": "EUR"
            },
            "created_at": "2024-11-20T01:30:38.039149Z",
            "updated_at": "2024-11-20T01:30:38.039149Z"
        },
        {
            "id": "f54db8c2-e541-41e8-b1af-6ea29d9beadb",
            "sku": "000002",
            "name": "BV Lean leather ankle boots",
            "category": "boots",
            "price": {
                "original": 990,
                "final": 693,
                "discount_percentage": "30%",
                "currency": "EUR"
            },
            "created_at": "2024-11-20T01:30:38.039149Z",
            "updated_at": "2024-11-20T01:30:38.039149Z"
        },
        {
            "id": "edb4bdf4-e7a1-40a7-8f19-1f0341f5eb71",
            "sku": "000003",
            "name": "Ashlington leather ankle boots",
            "category": "boots",
            "price": {
                "original": 710,
                "final": 497,
                "discount_percentage": "30%",
                "currency": "EUR"
            },
            "created_at": "2024-11-20T01:30:38.039149Z",
            "updated_at": "2024-11-20T01:30:38.039149Z"
        }
    ],
    "cursorId": "My00NTM0NTE="
}
```

### Folder structure
- Using Hexagonal Architecture to manage the product-service experience
```
/cmd
  /app
    - main.go             # Entry point for the application


/config                   # Application configuration
  config.go

/init-scripts             # Postgres db configuration, table and sample data insertion for docker
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

