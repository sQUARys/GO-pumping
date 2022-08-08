package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDb"

	dbInsertJSON = `INSERT INTO "order_table"( "order_id", "status", "store_id", "date_created") values($1 , $2 , $3 , $4)`
)

type Repository struct {
	DbStruct *sql.DB
}

func (repo *Repository) New() {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	repo.DbStruct = db

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

}

type Order struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	StoreId     int    `json:"store_id"`
	DateCreated string `json:"date_created"`
}

func (repo *Repository) Add(data Order) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.DbStruct.ExecContext(
		ctx,
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
