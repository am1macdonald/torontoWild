#!/bin/bash

# db_manage.sh
if [ -f .env ]; then 
  source .env
fi

# Database connection string
DB_CONN="${DB_URL}"

# Function to display usage information
usage() {
    echo "Usage: $0 <command>"
    echo "Commands:"
    echo "  up              - Run all pending migrations"
    echo "  down            - Rollback the last migration"
    echo "  reset           - Rollback all migrations and reapply them"
    echo "  status          - Show the status of migrations"
    echo "  generate        - Generate Go code from SQL using sqlc"
    echo "  create <name>   - Create a new migration file"
}

# Check if a command is provided
if [ $# -eq 0 ]; then
    usage
    exit 1
fi

# Execute command
case "$1" in
    up)
        echo "Running migrations up..."
        goose -dir ./sql/schema postgres "${DB_CONN}" up
        ;;
    down)
        echo "Rolling back the last migration..."
        goose -dir ./sql/schema postgres "${DB_CONN}" down
        ;;
    reset)
        echo "Resetting database..."
        goose -dir ./sql/schema postgres "${DB_CONN}" reset
        goose -dir ./sql/schema postgres "${DB_CONN}" up
        ;;
    status)
        echo "Migration status:"
        goose -dir ./sql/schema postgres "${DB_CONN}" status
        ;;
    generate)
        echo "Generating Go code from SQL..."
        sqlc generate
        ;;
    create)
        if [ $# -ne 2 ]; then
            echo "Error: Missing migration name"
            usage
            exit 1
        fi
        echo "Creating new migration '$2'..."
        goose -dir ./sql/schema create "$2" sql
        ;;
    *)
        echo "Error: Unknown command '$1'"
        usage
        exit 1
        ;;
esac

# If we've run migrations, also generate code
if [[ "$1" == "up" || "$1" == "down" || "$1" == "reset" ]]; then
    echo "Regenerating Go code..."
    sqlc generate
fi

echo "Done!"

