package main

import (
	"context"
	"dhlabs/backend/controller"
	"dhlabs/backend/database"
	"dhlabs/backend/services"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server             *gin.Engine
	transService       services.TransactService
	transacsController controller.TransController
	ctx                context.Context
	transactions       *mongo.Collection
	client             *mongo.Client
	err                error
	port               string
)

func init() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("error loading")
	// }

	// ctx = context.TODO()
	// peru := os.Getenv("MONGO_URI")
	// mongoconn := options.Client().ApplyURI(peru)
	// client, err = mongo.Connect(ctx, mongoconn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("mongodone")
	client, ctx := database.GetClient()

	transactions = client.Database("dhDB").Collection("transactions")
	transService = services.TransConstruct(transactions, ctx)
	transacsController = controller.TransControllerConstruct(transService)
	server = gin.Default()
	server.Use(cors.Default())

	port = os.Getenv("PORT")

}

func main() {

	defer client.Disconnect(ctx)
	basepath := server.Group("/chartapp")
	transacsController.Routes(basepath)

	log.Fatal(server.Run(":" + port))

}
