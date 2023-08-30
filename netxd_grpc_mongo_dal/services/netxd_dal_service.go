package services

import (
	"Netxd_grpc_mongo/netxd_grpc_mongo_dal/interfaces"
	"Netxd_grpc_mongo/netxd_grpc_mongo_dal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CustomerService handles customer-related operations.
type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

// InitCustomerService initializes a CustomerService instance.
func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomers {
	return &CustomerService{collection, ctx}
}

// CreateCustomer creates a new customer profile.
func (p *CustomerService) CreateCustomer(customer *models.Customers) (*models.CustomerResponse,error) {
	t := time.Now()
	customer.CreatedAt = t.String()
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = "true"

	// Insert customer data into the database.
	res, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}

	// Retrieve the newly created customer data from the database.
	newUser := &models.CustomerResponse{}
	query := bson.M{"_id": res.InsertedID}

	err = p.CustomerCollection.FindOne(p.ctx, query).Decode(newUser)
	if err != nil {
		return nil,err // Return nil to handle find error
	}

	return newUser,nil
}
