package service

import (
	"context"
	"fmt"
	"log"

	"github.com/parthverma/CRUDapplication/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin@cluster0.xaqu3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "Employee"
const collectionName = "EmployeeCollection"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Connection instance ready")
}

func InsertOneEmployee(employee model.Employee) string {
	inserted, err := collection.InsertOne(context.Background(), employee)
	if err != nil {
		log.Fatal(err)
	}

	return inserted.InsertedID.(primitive.ObjectID).String()
}

func GetAllEmployees() []primitive.M {

	filter := bson.D{{}}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var employees []primitive.M

	for cursor.Next(context.Background()) {
		var employee bson.M
		err := cursor.Decode(&employee)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
		defer cursor.Close(context.Background())
	}
	return employees
}

func UpdateOneEmployee(employee model.Employee, employeeID string) int64 {

	id, _ := primitive.ObjectIDFromHex(employeeID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": employee.Name, "age": employee.Age}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.ModifiedCount
}

func DeleteOneEmployee(employeeID string) int64 {
	id, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}

func DeleteAllEmployees() int64 {

	filter := bson.D{{}}

	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}
