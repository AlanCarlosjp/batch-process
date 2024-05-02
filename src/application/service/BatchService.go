package service

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"invoice/process/src/infrastructure/db"
	"invoice/process/src/model/entity"
	"log"
	"os"
	"time"
)

type BatchService struct {
	db *db.DbConnection
}

func NewBatchService() *BatchService {
	return &BatchService{
		db: db.NewDbConnectionBase(),
	}
}

func (b *BatchService) Migrate(path *string) (any, error) {
	file, err := os.Open(*path)
	if err != nil {
		log.Fatalln("Unable to read input file "+*path, err)
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse file as CSV for "+*path, err)
		return nil, err
	}
	orders := make([]entity.Order, 0)
	for _, record := range records {
		order := convertToOrder(record)
		orders = append(orders, order)
	}
	for _, order := range orders {
		fmt.Println(order)
	}

	return len(orders), nil
}

func (b *BatchService) CreateConnection(db_connection *string) {
	db.NewDbConnection(*db_connection)
}

func convertToOrder(record []string) entity.Order {
	order := entity.Order{
		ID:             uuid.New(),
		OrderDate:      time.Now(),
		InvoiceNumber:  record[0],
		CustomerID:     record[1],
		CustomerName:   record[2],
		ProductID:      record[3],
		ProductName:    record[4],
		Category:       record[5],
		QuantityBought: 0,
		SellingPrice:   0,
		UnitCost:       0,
		InvoiceSales:   0,
		InvoiceCost:    0,
		ShipmentDate:   time.Now(),
	}
	return order
}
