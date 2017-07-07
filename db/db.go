package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//Connect ...
func Connect(dbFileNamePtr *string) (*gorp.DbMap, error) {

	db, err := sql.Open("sqlite3", *dbFileNamePtr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
