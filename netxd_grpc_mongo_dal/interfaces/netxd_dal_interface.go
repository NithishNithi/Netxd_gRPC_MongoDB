package interfaces

import "Netxd_grpc_mongo/netxd_grpc_mongo_dal/models"





type ICustomers interface {
	CreateCustomer(customer *models.Customers) (*models.CustomerResponse, error)
}
