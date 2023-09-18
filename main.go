package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redmicelles/Hotel-Reservation/api"
	"github.com/redmicelles/Hotel-Reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017/"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ucoll := client.Database(dbname).Collection(userColl)
	user := types.User{
		FirstName: "Nathan",
		LastName:  "Daniel",
	}
	_, err = ucoll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var nathan types.User
	if err := ucoll.FindOne(ctx, bson.M{}).Decode(&nathan); err != nil {
		log.Fatal((err))
	}
	fmt.Println(nathan)
	listenAddr := flag.String("listenAddr", ":5000", "API Server Listening..")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleListUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
