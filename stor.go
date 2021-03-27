package stor

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseConnection interface {
	connectDB() (*sql.DB, error)
}

// DatabaseParams - parameters needed to create a new database connection
type DatabasePostgresParams struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	dbtype   string
	ssl      string
}

// DatabaseParams - parameters needed to create a new database connection
type DatabaseMySqlParams struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
	dbtype   string
	ssl      string
}

// DatabaseParams - parameters needed to create a new database connection
type DatabaseSqlite3Params struct {
	user     string
	password string
	dbname   string
	dbtype   string
}

func (d DatabasePostgresParams) connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", d.host, d.port, d.user, d.password, d.dbname, d.ssl))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d DatabaseMySqlParams) connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.user, d.password, d.host, d.port, d.dbname))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d DatabaseSqlite3Params) connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_auth&_auth_user=%s&_auth_pass=%s", d.dbname, d.user, d.password))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDatabase(d DatabaseConnection) (*sql.DB, error) {
	return d.connectDB()
}

func main() {
	fmt.Println("Stor: A golang package for exporting databases and uploading to AWS S3, FTP, Google Drive and saving locally")
	c := &DatabaseSqlite3Params{
		dbtype:   "sqlite3",
		dbname:   "mydb.db",
		user:     "",
		password: "",
	}
	conn, err := connectToDatabase(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(conn)
	defer conn.Close()
	//check if db connected successfully
	if err := conn.Ping(); err != nil {
		panic(err)
	}
	rows, err := conn.Query("")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
