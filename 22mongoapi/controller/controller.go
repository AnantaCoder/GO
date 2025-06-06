package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read from environment variables
	connectionString := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("COLLECTION_NAME")

	// Set client options
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// mongo db helpers - insert and other operations of the data .

func GetCollection() *mongo.Collection {
	return collection
}

// helpers written with small case since we are not exporting them outside this pkg
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db ", inserted.InsertedID)
}

// update a record
func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count", result.ModifiedCount)

}

// delete one and many movies in mongo db
func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("One movie is deleted : ", result)
}

// delete all movies in mongo

func deleteAllMovies() {
	filter := bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All movies deleted successfully, count: ", deleteResult.DeletedCount)
}

//all movies

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allmovies := getAllMovies()

	if len(allmovies) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No movies found",
		})
		return
	}

	//if 200ok and json encoded array
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allmovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
    

    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted all movies",
	})
}
