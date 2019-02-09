package data

import (
	"flag"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "Majid1212", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	server   = flag.String("server", "192.168.228.128", "the database server")
	user     = flag.String("user", "SA", "the database user")
	database = flag.String("database", "SimpleProject", "the database name")
)

var db *sqlx.DB

//GetDB ...
func GetDB() {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sqlx.Connect("sqlserver", connString)
	if err != nil {
		//log.Fatal("Open connection failed:", err.Error())
		panic(err)
	}

	//	fmt.Printf("Connected!\n")
	//	defer conn.Close()
	db = conn
	//return conn
}
