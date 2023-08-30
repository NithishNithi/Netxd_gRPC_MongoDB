package netxdcontrollers

import (
	"Netxd_gRPC_MongoDb/netxd_grpc_mongo_dal/interfaces"
	"Netxd_gRPC_MongoDb/netxd_grpc_mongo_dal/models"
	pro "Netxd_gRPC_MongoDb/netxd_grpc_mongo_proto/Customer_Protobuff"
	"context"
	"log"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomers
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerDetails) (*pro.CustomerResponse, error) {
	customer_record := &models.Customers{
		CustomerId: int(req.CustomerId),
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		BankId:     req.BankId,
		Balance:    float64(req.Balance),
	}

	// Call the CreateCustomer method of the CustomerService
	result, err := CustomerService.CreateCustomer(customer_record)
	if err != nil {
		log.Fatal(err)
	}

	// Create the responseCustomer only when there's no error
	responseCustomer := &pro.CustomerResponse{
		CustomerId: int32(result.CustomerId),
		CreatedAt:  result.CreatedAt,
	}

	return responseCustomer, nil
}
