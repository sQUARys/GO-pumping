package database

import (
	"RostPart4/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDb"

	dbInsertJSON = `INSERT INTO "order_table"( "order_id", "status", "store_id", "date_created") values($1 , $2 , $3 , $4)`
)

type LocalDB struct {
	DbStruct *sql.DB
}

func New() LocalDB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalln(err)
	}

	database := LocalDB{
		DbStruct: db,
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return database
}

func (d *LocalDB) Add(data models.Data) {
	_, err := d.DbStruct.Exec(
		dbInsertJSON,
		data.OrderId,
		data.Status,
		data.StoreId,
		data.DateCreated,
	)
	if err != nil {
		log.Print(err)
	}
}
