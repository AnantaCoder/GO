package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AEbhRhsd8iE51QDl
const connectionString = "mongodb+srv://gobhai:AEbhRhsd8iE51QDl@cluster0.dohk5tx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"


const dbName= "netflix"
const colName="watchlist"



//MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance reference 
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

