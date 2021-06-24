package database

import "fmt"

var (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "postgres"
	dbname    = "postgres"
	pgConnStr = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
)
