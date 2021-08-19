package DB

import(
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "abc123"
	dbname   = "provisions"
)

func ConnectServer() string{
	PSQLserver := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return PSQLserver
}