# TorontoWild: GTA Wildlife Mapping Platform

## Overview

TorontoBio is an open-source biodiversity mapping platform focused on Toronto and the Greater Toronto Area (GTA). It allows citizen scientists, researchers, and conservationists to log wildlife sightings, visualize biodiversity hotspots, and analyze species distribution patterns.

## Features

- User authentication via magic links
- Wildlife sighting logging with geolocation
- Interactive map visualization of sightings
- Species distribution modeling
- Biodiversity hotspot analysis
- Seasonal migration pattern visualization
- Data export for researchers

## Tech Stack

- Backend: Go
- Database: PostgreSQL with PostGIS extension
- In-memory Store: ValKey (Redis fork)
- API: RESTful, generated with sqlc
- Frontend: React with TypeScript (planned)
- Mapping: Leaflet or MapboxGL JS (planned)
- Data Visualization: D3.js (planned)

## Prerequisites

- Go 1.22.*
- PostgreSQL 13+ with PostGIS extension
- ValKey (Redis fork)
- sqlc
- goose (for database migrations)

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/torontobio.git
   cd torontobio
   ```

2. Set up the database:
   ```
   createdb torontobio
   psql -d torontobio -c "CREATE EXTENSION postgis;"
   ```

3. Run database migrations:
   ```
   goose postgres "user=yourusername dbname=torontobio sslmode=disable" up
   ```

4. Generate Go code from SQL:
   ```
   sqlc generate
   ```

5. Set up environment variables:
   ```
   export DB_HOST=localhost
   export DB_USER=yourusername
   export DB_NAME=torontobio
   export DB_PORT=5432
   export VALKEY_ADDR=localhost:6379
   ```

6. Run the application:
   ```
   go run cmd/torontobio/main.go
   ```


## Database Management

We use a custom shell script to manage database migrations and code generation. The script provides the following commands:

- `./db_manage.sh up`: Run all pending migrations
- `./db_manage.sh down`: Rollback the last migration
- `./db_manage.sh reset`: Rollback all migrations and reapply them
- `./db_manage.sh status`: Show the status of migrations
- `./db_manage.sh generate`: Generate Go code from SQL using sqlc
- `./db_manage.sh create <name>`: Create a new migration file

To use the script:

1. Ensure your environment variables are set:
   ```
   export DB_HOST=localhost
   export DB_USER=yourusername
   export DB_NAME=torontobio
   ```

2. Run the desired command, for example:
   ```
   ./db_manage.sh up
   ```

This script will automatically regenerate Go code using sqlc after running migrations.

## API Endpoints

- `POST /request-magic-link`: Request a magic link for authentication
- `GET /login/{token}`: Verify magic link and log in
- `POST /sightings`: Log a new wildlife sighting
- `GET /sightings`: Retrieve wildlife sightings
- `GET /hotspots`: Get biodiversity hotspots
- `GET /species-distribution`: Get species distribution data

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- OpenStreetMap for base map data
- iNaturalist for initial species data
- Government of Ontario for environmental data

```
