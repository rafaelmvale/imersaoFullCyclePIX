package main

import (
	"os"

	"github.com/jinzhu/gorm"

	"github.com/rafaelmvale/imersaoFullCyclePIX/application/grpc"
	"github.com/rafaelmvale/imersaoFullCyclePIX/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
