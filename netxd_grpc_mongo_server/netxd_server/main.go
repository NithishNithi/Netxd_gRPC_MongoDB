package main

import (
	"Netxd_gRPC_MongoDb/netxd_grpc_mongo_dal/services"
	pro "Netxd_gRPC_MongoDb/netxd_grpc_mongo_proto/Customer_Protobuff"
	netxdconfig "Netxd_gRPC_MongoDb/netxd_grpc_mongo_server/netxd_config"
	netxdconstants "Netxd_gRPC_MongoDb/netxd_grpc_mongo_server/netxd_constants"
	netxdcontrollers "Netxd_gRPC_MongoDb/netxd_grpc_mongo_server/netxd_controllers"
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	profileCollection := netxdconfig.GetCollection(client, "Netxd_Bank", "Customers")
	netxdcontrollers.CustomerService = services.InitCustomerService(profileCollection, context.Background())
}

func main() {
	mongoclient, err := netxdconfig.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxdconstants.Port)

	if err != nil {
		fmt.Println("Failed TO Listen", err)
	}
	s := grpc.NewServer()

	pro.RegisterCustomerServiceServer(s, &netxdcontrollers.RPCServer{})

	fmt.Println("Server listening on", netxdconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
