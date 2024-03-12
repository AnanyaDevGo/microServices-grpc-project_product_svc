package main

import (
	"fmt"
	"log"
	"net"

	// "github.com/YOUR_USERNAME/go-grpc-product-svc/pkg/config"
	// "github.com/YOUR_USERNAME/go-grpc-product-svc/pkg/db"
	// pb "github.com/YOUR_USERNAME/go-grpc-product-svc/pkg/pb"
	// services "github.com/YOUR_USERNAME/go-grpc-product-svc/pkg/services"
	"github.com/AnanyaDevGo/microServices-grpc-project_product_svc/pkg/config"
	"github.com/AnanyaDevGo/microServices-grpc-project_product_svc/pkg/db"
	"github.com/AnanyaDevGo/microServices-grpc-project_product_svc/pkg/pb"
	"github.com/AnanyaDevGo/microServices-grpc-project_product_svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    h := db.Init(c.DBUrl)

    lis, err := net.Listen("tcp", c.Port)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    fmt.Println("Product Svc on", c.Port)

    s := services.Server{
        H: h,
    }

    grpcServer := grpc.NewServer()

    pb.RegisterProductServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalln("Failed to serve:", err)
    }
}