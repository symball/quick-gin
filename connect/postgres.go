package connect

import (
	"log"
	"os"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

func Postgres() sqlbuilder.Database {

	DbHost := LookupEnvOrExit("DB_HOST")
	DbSchema := LookupEnvOrExit("DB_SCHEMA")
	DbUser := LookupEnvOrExit("DB_USER")
	DbPass := LookupEnvOrExit("DB_PASS")

	var settings = postgresql.ConnectionURL{
		Host:     DbHost,
		Database: DbSchema,
		User:     DbUser,
		Password: DbPass,
	}

	db, err := postgresql.Open(settings)
	if err != nil {
		log.Print("Could not connect to DB. Please check the database parameters for any errors.")
		os.Exit(1)
	}
	log.Printf("Connected to DB at host: %s & database: %s, using user: %s", DbHost, DbSchema, DbUser)
	return db
}
