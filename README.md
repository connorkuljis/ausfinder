# Ausfinder

A Go-based web application for searching and retrieving Australian business information.

## Dataset

[https://data.gov.au/data/dataset/bc515135-4bb6-4d50-957a-3713709a76d3/resource/55ad4b1c-5eeb-44ea-8b29-d410da431be3/download/business_names_202503.csv](https://data.gov.au/data/dataset/bc515134-4bb6-4d50-957a-3713709a76d3/resource/55ad4b1c-5eeb-44ea-8b29-d410da431be3/download/business_names_202503.csv)

## Overview

Backtrace provides a simple web interface to search for Australian businesses by name or ABN. It utilizes the Australian Business Register (ABR) API and maintains a local SQLite database for efficient queries.

## Features

- Search businesses by name
- Filter by state
- View detailed business information
- Direct link to LinkedIn company search

## Installation

```bash
# Clone the repository
git clone https://github.com/connorkuljis/ausfinder.git

# Navigate to the project directory
cd ausfinder

# Install dependencies
go mod tidy

# Build the application
go build -o main

# Run the service
./main
```

## Requirements

- Go 1.18+
- SQLite3

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | / | Redirects to /search |
| GET    | /search | Search page with optional query parameters |
| GET    | /search/business/:id | Detailed view of a specific business by ABN |

## API Usage

### GET /search

Searches for businesses with optional filtering.

**Query Parameters:**
- `q`: Search term for business name
- `state`: Filter by state (optional)

**Example:**
```
/search?q=coffee&state=NSW
```

**Response:**
Renders a page with business search results.

### GET /search/business/:id

Retrieves detailed information about a specific business.

**Parameters:**
- `id`: The ABN of the business

**Example:**
```
/search/business/12345678901
```

**Response:**
Renders a page with detailed business information and a LinkedIn search link.

## Database

The application uses SQLite with WAL journal mode for data storage. The database should be located at `db/db.sqlite3`.

Tables:
- `business_search`: Used for full-text search of business names
- `business_names`: Contains detailed business information

## Project Structure

```
backtrace/
├── internal/
│   ├── model/       # Data models
│   └── renderer/    # HTML template renderer
├── public/          # Static assets
├── templates/       # HTML templates
└── db/              # SQLite database
```

## External Service Integration

The application integrates with the Australian Business Register (ABR) XML Search API for data validation and retrieval.

## Configuration

The server runs on port 8080 by default.

## License

MIT
