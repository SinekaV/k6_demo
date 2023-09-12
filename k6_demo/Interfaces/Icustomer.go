package interfaces

import (
	"k6_demo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface {
	CreateCustomer(*models.Customer)(*mongo.InsertOneResult,error)
}