# Database Setup Guide

This project is configured to use PostgreSQL with Railway.

## Railway PostgreSQL Database

Your Railway database is hosted at: `postgres-production-926b.up.railway.app`

## Setup Instructions

### Option 1: Using DATABASE_URL (Recommended)

1. Get your connection string from Railway:
   - Go to your Railway project dashboard
   - Navigate to your PostgreSQL service
   - Click on "Connect" or "Variables"
   - Copy the `DATABASE_URL` or `POSTGRES_URL`

2. Create a `.env` file in the project root (if using dotenv gem) or set environment variables:
   ```bash
   DATABASE_URL=postgresql://user:password@postgres-production-926b.up.railway.app:5432/railway
   ```

### Option 2: Using Individual Environment Variables

If you prefer to use individual variables, create a `.env` file with:

```bash
DB_HOST=postgres-production-926b.up.railway.app
DB_PORT=5432
DB_NAME=railway
DB_USER=postgres
DB_PASSWORD=KAeixPsAJvutpccGkLZEOcjRJqEJsqnz
```

## Running Migrations

After setting up your environment variables, run:

```bash
# Create the database (if needed)
rails db:create

# Run migrations
rails db:migrate

# Seed the database (optional)
rails db:seed
```

## Environment Variables

The application supports the following environment variables:

- `DATABASE_URL` - Full PostgreSQL connection string (takes precedence)
- `DB_HOST` - Database hostname
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `RAILS_MAX_THREADS` - Maximum database connections (default: 5)
- `RAILS_LOG_LEVEL` - Log level (default: info)

## Notes

- The `.env` file is already in `.gitignore` and won't be committed to version control
- Make sure you have the `pg` gem installed (already in Gemfile)
- For production, set these variables in your deployment platform (Railway, Heroku, etc.)

