# README

This README would normally document whatever steps are necessary to get the
application up and running.

## Database Setup

This project uses PostgreSQL with Railway. See [DATABASE_SETUP.md](DATABASE_SETUP.md) for detailed setup instructions.

Quick setup:
1. Set your `DATABASE_URL` environment variable with your Railway PostgreSQL connection string
2. Run `rails db:migrate` to set up the database schema

## Configuration

* Ruby version: See `.ruby-version` or `Gemfile`
* Database: PostgreSQL (configured for Railway)
* See `DATABASE_SETUP.md` for database configuration details

## Database Creation and Initialization

```bash
rails db:create
rails db:migrate
rails db:seed
```

## How to run the test suite

```bash
rspec
```

## Services (job queues, cache servers, search engines, etc.)

* Solid Queue for background jobs
* Solid Cache for caching
* Solid Cable for Action Cable

## Deployment instructions

* Configure `DATABASE_URL` in your deployment platform
* The application is configured to use PostgreSQL in production
