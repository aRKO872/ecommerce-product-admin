# ecommerce-product-admin

This is a simple E-Commerce Product Admin Panel BE, which is distributed consisting of total 6 microservices operating over HTTP and GRPC calls. Writes to all product management microservices and logging microservice are done via Message Queues (Kafka)

## Table Of Contents
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Development](#development)
- [Project Structure](#project-structure)
    - [Explanation of Services](#explanation-of-services)
    - [Explanation of Directories for all Services](#explanation-of-directories-for-all-services)

## Prerequisites
Before setting up the project, ensure you have the following installed:
- [Go 1.24.4+](https://golang.org/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)

## Installation
1. **Clone the repository**:
```bash
  git clone https://github.com/aRKO872/ecommerce-product-admin.git
```

2. **Install dependencies**: cd into each service and ensure that all third party packages are installed and proper GO build is being generated.
```bash
  cd backend-facing-frontend
  go mod tidy
  go build
```

3. **Setup environment variables**: cd into each service and run the following commands to create .env file for that particular service. You could put in your own configurations as you see fit.
```bash
  cp .env.example .env
```

Example `.env` file:
```ini
  APP_ID="core-engine"
  PORT="7001"
  HOST="0.0.0.0"
  INVENTORY_MSC_ADDR="inventory:7003"
  ORDERS_MSC_ADDR="orders:7004"
  PRODUCTS_MSC_ADDR="products:7005"
```

## Development

- To build (if not already built) and run all microservice instances and set up containers for MySQL DB, Kafka and GoDoc (Documentation)  :
```bash
  make up
```

- To bring down all containers (if currently running) :
```bash
  make down
```

- To check logs for a certain service : 
```bash
  make logs SERVICE=<service-name>
```

- To check active/running containers : 
```bash
  make ps
```

- To check all containers (active/inactive) : 
```bash
  make ps-all
```

- To clear all containers (active/inactive), with their built up images and volumes :
```bash
  make clean
```

- Below are the default godoc ports to access documentation for respective services after all containers are successfully run and are active :  
    - backend-facing-frontend : `localhost:6060/`

## Project Structure

```bash
backend-facing-frontend
    ├── .env
    ├── Dockerfile
    ├── client
        ├── client.go
        └── heartbeat.go
    ├── controllers
        ├── controller.go
        ├── heartbeat.go
        ├── literals.go
        └── router.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
        ├── env.go
        └── models.go
    ├── prepare
        └── prepare.go
    └── services
        ├── heartbeat.go
        └── service.go
core-engine
    ├── .env
    ├── Dockerfile
    ├── client
        ├── client.go
        └── heartbeat.go
    ├── controllers
        ├── controller.go
        ├── heartbeat.go
        └── router.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
        └── env.go
    ├── prepare
        └── prepare.go
    └── services
        ├── heartbeat.go
        └── service.go
inventory-msc
    ├── .env
    ├── Dockerfile
    ├── controllers
        ├── controller.go
        ├── heartbeat.go
        └── router.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
        └── env.go
    ├── prepare
        └── prepare.go
    └── services
        ├── heartbeat.go
        └── service.go
logger-msc
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── prepare
        └── prepare.go
orders-msc
    ├── .env
    ├── Dockerfile
    ├── controllers
        ├── controller.go
        ├── heartbeat.go
        └── router.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
        └── env.go
    ├── prepare
        └── prepare.go
    └── services
        ├── heartbeat.go
        └── service.go
products-msc
    ├── .env
    ├── Dockerfile
    ├── [controllers]
        ├── controller.go
        ├── heartbeat.go
        └── router.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
        └── env.go
    ├── prepare
        └── prepare.go
    └── services
        ├── heartbeat.go
        └── service.go
Makefile
README.md
docker-compose.yml
```

### Explanation of Services
1. **backend-facing-frontend**  
This is the only service in the backend, which would be exposed to the Frontend. Being a HTTP Service, it's sole purpose is to relay all the calls from FE to the Core-Engine GRPC Service, which would further process the request and provide an expected response.  
List of HTTP routes : 
    - /bff-service/health (GET) :  
        This endpoint makes sure all underlying GRPC microservices are up and running. We get success response if all microservices are active.
        - Response :
            - Success :
                ```ini
                {
                    "status": "OK"
                }
                ```
            - Failure :
                ```ini
                {
                    "error": "error message"
                }
                ```

2. **core-engine**  
This is the intermediary GRPC service between the Backend Facing Frontend HTTP service and the Inventory, Products and Orders GRPC services which communicate with the Database for persistence and fetching of data. As per the use case of each HTTP endpoint andtheir resepective business logic, calls are relayed to other GRPC services for fetching and persistence of data. Kafka calls are made for asynchronous persistence of data.

3. **logger-msc**  
A GRPC service which is communicated with via Kafka pub-sub, for logging. The Core-Engine, Inventory, Products and Orders GRPC services send logs to the logging topic that is consumed in turn by this service, in order to print logs in it's terminal, for now. Future iterations might involve received logs being fed to third party monitoring services like AWS Cloudwatch or GCP. 

4. **inventory-msc**  
GRPC service that involves business logic for management and fetching of product inventories in the Database.

5. **orders-msc**  
GRPC service that involves business logic for management and fetching of user created orders for products in the Database.

6. **products-msc**  
GRPC service that involves business logic for management and fetching of product information in the Database.


### Explanation of Directories for all Services
1. **prepare**  
Consists of prepare.go, which initializes and bundles together the core components required for any service to function. It is responsible for reading and validating environment config, setting up gRPC connections to other microservices, and preparing the HTTP/GRPC routing layer for incoming requests. This package serves as the entry point
for bootstrapping the service runtime.

2. **controllers**  
Consists of the primary Controller object, its Router function and all endpoint level controllers which are to be mapped to the ServiceRouter object that would be in turn used to set up our HTTP/GRPC server.  
The controller consists of the Service object which holds the primary business logic for all endpoints. Th endpoint level controllers are responsible for extracting and validating requests and sending back responses received from corresponding Service object functions.

3. **services**  
Consists of the Service Object which holds the Client and DAO objects for inter-service communication and database operations respectively. It consists of all the business logic for a certain endpoint.

4. **client**  
Consists of the Client object which holds the GRPC Clients of other GRPC microservices and Kafka connector for async writes. This object is solely responsible for making inter-service calls (both synchronous GRPC calls and asynchronous Kafka calls).

5. **dao**  
Consists of the DAO object which holds the singleton instance of the Database Connector (MySQL). Attached to this object are the functionalities for updating or fetching data from various tables in the Database as per the business logic in the service layer. This directory is only present in those GRPC services which communicate with the Database.

6. **models**  
Consists of all the models (structs/interfaces) to be used by functions in all directories of the service.

