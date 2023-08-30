package main

import (
	pro "Netxd_gRPC_MongoDb/netxd_grpc_mongo_proto/Customer_Protobuff"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to Connect", err)
	}
	defer conn.Close()
	client := pro.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pro.CustomerDetails{CustomerId: 101, FirstName: "Nithish", LastName: "T", BankId: "1664Bank1", Balance: 55000.00})
	if err != nil {
		log.Fatalf("Failed to call CreateCustomer: %v", err)
	}

	fmt.Println("Response:", response.CustomerId, response.CreatedAt)
}
