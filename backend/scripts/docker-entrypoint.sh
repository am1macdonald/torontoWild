#!/bin/sh
set -e

# Run database migrations
./scripts/db_manage.sh up

# Start the main application
exec air
