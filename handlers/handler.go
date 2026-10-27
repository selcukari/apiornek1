package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"web-api/config"
	"web-api/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all users
func GetAllContents(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("newsPopulerDb").Collection("Contents")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var contents []models.Content
	for cursor.Next(context.Background()) {
		var content models.Content
		if err := cursor.Decode(&content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		uuidValue, err := uuid.FromBytes(content.ID.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		content.Id = uuidValue.String()
		contents = append(contents, content)
	}

	json.NewEncoder(w).Encode(contents)
}

// Get a user by ID
func GetContentID(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	id := strings.TrimPrefix(r.URL.Path, "/contents/")

	// String değerini UUID'ye dönüştür
	uuidValue, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, "Geçersiz UUID", http.StatusBadRequest)
		return
	}

	// UUID'yi BSON Binary subtype 4 biçimine dönüştür
	uuidBinary := primitive.Binary{
		Subtype: 4,
		Data:    uuidValue[:],
	}

	collection := client.Database("newsPopulerDb").Collection("Contents")
	var content models.Content
	if err := collection.FindOne(context.Background(), bson.M{"_id": uuidBinary}).Decode(&content); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	getUuidValue, err := uuid.FromBytes(content.ID.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content.Id = getUuidValue.String()

	json.NewEncoder(w).Encode(content)
}
