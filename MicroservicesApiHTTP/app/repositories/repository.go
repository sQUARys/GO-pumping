package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sQUARys/GO-pumping/MicroservicesApiHTTP/app/model"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myUser"
	password = "myPassword"
	dbname   = "myDb"

	connectionFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"

	dbOrdersByIdRequest = "SELECT * FROM order_table WHERE order_id = $1"
)

type Repository struct {
	DbStruct *sql.DB
}

func New() *Repository {
	stringConnection := fmt.Sprintf(connectionFormat, host, port, user, password, dbname)

	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatalln(err)
	}

	repo := Repository{DbStruct: db}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &repo
}

func (repo *Repository) GetOrdersById(id int) (model.Order, error) {
	order := model.Order{}
	err := repo.DbStruct.QueryRow(dbOrdersByIdRequest, id).Scan(&order.OrderId, &order.Status, &order.StoreId, &order.DateCreated)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
