package main

import (
	"fmt"

	"grpc-prj/api/articlepb"
	"grpc-prj/api/repository"
	"grpc-prj/api/service"
	"grpc-prj/config/database"
	"grpc-prj/internal/crawls"
	"log"
	"net"

	"google.golang.org/grpc"
)

func init() {
	database.ConnectDB()
	crawls.Crawl()
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50069")
	if err != nil {
		log.Fatal(err)
	}

	articleRepo := repository.NewArticleRepository(database.DBConn)
	server := service.Newserver(articleRepo)
	s := grpc.NewServer()
	articlepb.RegisterArticleServiceServer(s, server)
	fmt.Println("Server is running ...")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

