package entity

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID             uuid.UUID
	OrderDate      time.Time
	InvoiceNumber  string
	CustomerID     string
	CustomerName   string
	ProductID      string
	ProductName    string
	Category       string
	QuantityBought int
	SellingPrice   float64
	UnitCost       float64
	InvoiceSales   float64
	InvoiceCost    float64
	ShipmentDate   time.Time
}
