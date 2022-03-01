package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"go/printer"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://<username>:<password>@cluster0.wd7ju.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongoDB
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect with mongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	collection = client.DataBase(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection is ready")

}

//mongo Db Helpers

//insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted One Movie in DB with Id: ", inserted.InsertedID)

}

//update one record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M("$set": bson.M{"watched": true}) 

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)

}

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	fiter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: " , deleteCount)

}

//delete all records from mongoDb
func deleteAllMovies() int64{
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted", deleteResult.deleteCount)
	return deleteResult.deleteCount


}


//get all movies from db

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil(){
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M
		err := cur.Decode(&movie)
		if err!= nil{
			log.Fatal(err)
		}
		movie = append(movies, movie)

	}

	defer cur.Close(context.Background())
	return movies
}


//actual controller file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}


fucn CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)




}


func MarkAsWatched(w http.ResponseWriter, r *http.Reques){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params, _ := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}


func deleteMovie(w http.ResponseWriter, r *http.Reques){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params, _ := mux.Vars(r)
	count :=  deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}