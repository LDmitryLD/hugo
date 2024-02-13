package db

import (
	"database/sql"
	"fmt"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewSqlDB(dbConf config.DB) (*sqlx.DB, *adapter.SQLAdapter, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name)
	var dbRaw *sql.DB

	dbRaw, err := sql.Open(dbConf.Driver, dsn)
	if err != nil {
		return nil, nil, err
	}
	err = dbRaw.Ping()
	if err != nil {
		return nil, nil, err
	}

	db := sqlx.NewDb(dbRaw, dbConf.Driver)
	sqlAdapter := adapter.NewSQLAdapter(db)

	return db, sqlAdapter, nil
}
