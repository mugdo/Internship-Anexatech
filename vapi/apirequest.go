package vapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	database "quickstart/dataBase"
	"quickstart/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Getproducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Database connection
	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)
	//Find all data from database
	productCollection := database.Collection("product")
	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var productSum []model.Product
	if err = cursor.All(ctx, &productSum); err != nil {
		log.Fatal(productSum)
	}
	// Error Message for Empty Database
	if productSum == nil {
		msg := struct {
			Error string `json:"error"`
		}{
			Error: "No Product Found into the DataBase. Insert one",
		}
		json.NewEncoder(w).Encode(msg)
	} else {
		json.NewEncoder(w).Encode(productSum)
	}

}

func Getproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// id value from api and transform  primitive hex value convert(string to primitive.ObjectID)
	prams := mux.Vars(r)
	oid, err := primitive.ObjectIDFromHex(prams["id"])
	if err != nil {
		log.Fatal(err)
	}
	//Database connection
	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)

	productCollection := database.Collection("product")

	var collectProduct model.Product
	if err = productCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&collectProduct); err != nil {
		fmt.Println(err)
	}
	// If it dosen't get any data it give a Error message
	if collectProduct.ID.Hex() == "000000000000000000000000" {
		msg := struct {
			Error string `json:"error"`
		}{
			Error: "No Product Found",
		}
		json.NewEncoder(w).Encode(msg)
	} else {
		json.NewEncoder(w).Encode(collectProduct)
	}
}
func Addproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//grab data from postman
	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		log.Fatal(err)
	}
	//Database connection
	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)
	//Insert data into database
	productCollection := database.Collection("product")
	insertResult, inserterr := productCollection.InsertOne(ctx, newProduct)
	if inserterr != nil {
		log.Fatal(inserterr)
	}
	json.NewEncoder(w).Encode(insertResult)

}
func Deleteproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//id  convert(string to primitive.ObjectID)
	prams := mux.Vars(r)
	oid, err := primitive.ObjectIDFromHex(prams["id"])
	if err != nil {
		log.Fatal(err)
	}
	//Database connection
	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)

	productCollection := database.Collection("product")
	var productConstroctor model.Product
	//Delet product by using given id
	if err = productCollection.FindOneAndDelete(ctx, bson.M{"_id": oid}).Decode(&productConstroctor); err != nil {
		log.Fatal(err)
	}
	if productConstroctor.ID.Hex() == "000000000000000000000000" {
		msg := struct {
			Error string `json:"error"`
		}{
			Error: "No Product Found into the DataBase. Insert one",
		}
		json.NewEncoder(w).Encode(msg)
	} else {
		json.NewEncoder(w).Encode(productConstroctor)
	}
	json.NewEncoder(w).Encode(productConstroctor)

}
func Updateproduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//id  convert(string to primitive.ObjectID)
	prams := mux.Vars(r)
	oid, _ := primitive.ObjectIDFromHex(prams["id"])

	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)

	productCollection := database.Collection("product")

	var reqProduct model.Product
	var dbProduct model.Product
	//grab data from postman
	err := json.NewDecoder(r.Body).Decode(&reqProduct)
	if err != nil {
		log.Fatal(err)
	}
	err = productCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&dbProduct)
	if err != nil {
		log.Fatal(err)
	}

	// assign data into second valu for compare empy value
	Code1 := reqProduct.Code
	Name1 := reqProduct.Name
	Description1 := reqProduct.Description
	Price1 := reqProduct.Price
	Count1 := reqProduct.Count
	Discount1 := reqProduct.Discount
	var Sizes1 []string = reqProduct.Sizes
	var Colors1 []string = reqProduct.Colors
	var Category1 []string = reqProduct.Category

	// Assign valu from database if it get nul value
	if reqProduct.Code == "" {
		Code1 = dbProduct.Code
	}
	if reqProduct.Name == "" {
		Name1 = dbProduct.Name
	}
	if reqProduct.Count == 0 {
		Count1 = dbProduct.Count
	}
	if reqProduct.Price == 0 {
		Price1 = dbProduct.Price
	}
	if reqProduct.Description == "" {
		Description1 = dbProduct.Description
	}
	if reqProduct.Sizes == nil {
		Sizes1 = dbProduct.Sizes
	}
	if reqProduct.Colors == nil {
		Colors1 = dbProduct.Colors
	}
	if reqProduct.Discount == "" {
		Discount1 = dbProduct.Discount
	}
	if reqProduct.Category == nil {
		Category1 = dbProduct.Category
	}
	// update new value
	updateField := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "code", Value: Code1},
			{Key: "name", Value: Name1},
			{Key: "description", Value: Description1},
			{Key: "price", Value: Price1},
			{Key: "count", Value: Count1},
			{Key: "discount", Value: Discount1},
			{Key: "sizes", Value: Sizes1},
			{Key: "colors", Value: Colors1},
			{Key: "category", Value: Category1},
		}},
	}
	if _, err = productCollection.UpdateOne(ctx, bson.M{"_id": oid}, updateField); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(updateField)

}
func Productcategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	prams := mux.Vars(r)
	oid := prams["value"]

	database, ctx, cancel := database.CreateConnection()
	defer cancel()
	defer database.Client().Disconnect(ctx)

	productCollection := database.Collection("product")
	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var addCategory []model.Product
	var temp model.Product
	// create a value into which the single document can be decoded
	for cursor.Next(ctx) {

		err := cursor.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		// search valu into category if get then added
		for _, value := range temp.Category {
			if value == oid {
				addCategory = append(addCategory, temp)
			}

		}

	}
	//Response back
	json.NewEncoder(w).Encode(addCategory)
}
