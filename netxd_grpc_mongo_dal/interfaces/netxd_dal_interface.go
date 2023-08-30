package interfaces

import "Netxd_gRPC_MongoDb/netxd_grpc_mongo_dal/models"

type ICustomers interface {
	CreateCustomer(customer *models.Customers) (*models.CustomerResponse, error)
}
