# Software Engineering Challenge

## Overview

This application calculates the number of packs that needed to ship to customer.

## Requirements
- Go 1.20 or higher
- Docker (Optional)

## Quick Start
### Clone the repository
```bash
git clone https://github.com/avag-sargsyan/testgs.git
cd testgs
``` 

### Building the application
Using Go:
```bash
make build
```

Using Docker:
```bash
make docker-build
```

### Running the application
Using Go:
```bash
make run
```

Using Docker:
```bash
make docker-run
```

### Running the tests
```bash
make test
```

### Cleaning up
```bash
make clean
```

## How to use
### API Endpoints
The application serves an HTTP API on port 8080 (which can be configured).
- GET /api/packs?order={order}: Calculates packs for the given order size.

## Configuration
The application can be configured using environment variables. The following environment variables are supported:
- GS_HTTP_ADDRESS: The port on which the HTTP server will listen. Defaults to 8080.
- GS_PACK_SIZES: The pack sizes that are available. Defaults to 250, 500, 1000, 2000, 5000.

## Deployment
The application can be deployed as a standalone binary or within a Docker container.
