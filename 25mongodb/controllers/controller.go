package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "mongodb/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// paste the connection string over here
const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

// Most imp

var collection *mongo.Collection

// connect with mongoDB
// init -> it is a special method that just once for the first itme
func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection is success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready")
}

// mongodb helper
// insert one record
func insertOneMovie(movie model.NetFlix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted one movie with id: ", inserted.InsertedID)
}

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.ModifiedCount)
}

// delete one record
func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
}

// delete all content
func deleteAllMovie() {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
}

// get all movie from db
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.TODO()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	fmt.Println(movies)
	return movies
}

// actual controller - file
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www/form-urlencode")
	movies := getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www/form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.NetFlix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www/form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/x-www/form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www/form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovie()
	json.NewEncoder(w).Encode("All movies deleted")
}