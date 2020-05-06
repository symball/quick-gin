# Quick Gin

This is a small Gin-gonic based project to be used as part of a coding challenge. Precompiled artifacts for both the [migration tool](https://github.com/golang-migrate/migrate) and project are provided in the root directory

## Quickstart

Run the `dev.sh` script. This will:

* Check whether you have docker installed
* Check open ports
* Fire up the postgresql support service
* Run database migrations
* Startup the gin service

## Routes

There are only two `GET` routes for this API

`/` - A simple healthcheck
`/projects` - Retrieves a JSON array of projects within the DB

## Notes

To run the unit tests, run `go test ./...`
