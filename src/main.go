package main

import (
	"flag"
	"net/http"
	"os"
	"thebombanation.com/website/pkg/models/mysql"
	"database/sql"
	"html/template"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

)

var tpl *template.Template

type application struct {
	infoLog *log.Logger
	errLog *log.Logger
	DB *mysql.UserModel
}

//keep essentials in main
func main() {
	addr := flag.String("address", ":8080", "Default port")
	dsn := flag.String("dns", "mob:another@/practice?parseTime=true", "MySQL data source name")
	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errLog := log.New(os.Stderr, "ERROR\t", log.LstdFlags | log.Lshortfile)
	flag.Parse()
	flag.Parsed()
	db, err:= openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog: infoLog,
		errLog: errLog,
		DB: &mysql.UserModel{DB: db},
	}
	//routes() handles all paths and returns a server mux
	mux := app.routes()

	srv := &http.Server{
		Addr: *addr,
		Handler: mux,
		ErrorLog: errLog,
		WriteTimeout: 15 * time.Second,
		ReadHeaderTimeout:  15 * time.Second,
	}

	infoLog.Printf("Starting server on port %v", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
 // opening and testing DB connection
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

