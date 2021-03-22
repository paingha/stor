package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DatabaseParams - parameters needed to create a new database connection
type DatabaseParams struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	dbtype   string
	ssl      string
}

func createDbConnectionString(databaseInput DatabaseParams) map[string]string {
	switch databaseInput.dbtype {
	case "postgres":
		return map[string]string{
			"databaseType":     "postgres",
			"connectionString": fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseInput.host, databaseInput.port, databaseInput.user, databaseInput.password, databaseInput.dbname),
		}
	case "mysql":
		return map[string]string{
			"databaseType":     "mysql",
			"connectionString": fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", databaseInput.user, databaseInput.password, databaseInput.host, databaseInput.port, databaseInput.dbname),
		}
	default:
		return map[string]string{
			"databaseType":     "postgres",
			"connectionString": fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseInput.host, databaseInput.port, databaseInput.user, databaseInput.password, databaseInput.dbname),
		}
	}
}

func connectToDatabase(databaseInput DatabaseParams) {
	db, err := sql.Open(createDbConnectionString(databaseInput)["databaseType"], createDbConnectionString(databaseInput)["connectionString"])
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//check if db connected successfully
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Stor: A golang package for exporting databases and uploading to AWS S3, FTP, Google Drive and saving locally")
}
