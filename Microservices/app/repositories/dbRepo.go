package repositories

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sQUARys/GO-pumping/app/model"
	"log"
	"strings"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDb"

	dbInsertJSON = `INSERT INTO "order_table"( "order_id", "status", "store_id", "date_created") VALUES `

	format = "(%d , '%s' , %d , '%s'),"
)

type Repository struct {
	DbStruct *sql.DB
}

func New() *Repository {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	repo := Repository{
		DbStruct: db,
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &repo
}

func (repo *Repository) Add(orders []model.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var (
		formattedOrders []string
		dbInsertRequest string
	)

	for i := 0; i < len(orders); i++ {
		formattedOrders = append(formattedOrders, fmt.Sprintf(format, orders[i].OrderId, orders[i].Status, orders[i].StoreId, orders[i].DateCreated))
	}

	dbInsertRequest = strings.Join(formattedOrders, "")
	dbInsertRequest = strings.TrimSuffix(dbInsertRequest, ",")
	dbInsertRequest = dbInsertJSON + dbInsertRequest

	_, err := repo.DbStruct.ExecContext(
		ctx,
		dbInsertRequest,
	)
	if err != nil {
		return err
	}
	return nil
}
